package middleware

import (
	"github.com/gin-gonic/gin"
	"seat-service/response"
	"seat-service/utils"
	"strconv"
	"strings"
	"time"
)

var resp response.CustomResponse

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			resp.FailBasedCode(response.CodeIllegalLogin, gin.H{"reload": true})
			c.Abort()
			return
		}
		// 按空格分割
		parts := strings.SplitN(token, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Token") {
			resp.FailBasedCode(response.CodeLoginFailure, gin.H{"reload": true})
			c.Abort()
			return
		}
		jwt := utils.NewJWT()
		mc, err := jwt.ParseToken(parts[1])
		if err != nil && err.Error() != "Token is expired" {
			resp.FailBasedCode(response.CodeTokenInvalid, gin.H{"reload": true})
			c.Abort()
			return
		}

		//续签
		if mc.ExpiresAt.Unix()-time.Now().Unix() < time.Now().Add(time.Hour*time.Duration(mc.Buffer)).Unix() || err.Error() == "Token is expired" {
			newToken, _ := jwt.CreateTokenByOldToken(token, mc.UserClaims)
			newClaims, _ := jwt.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Unix(), 10))
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("userClaims", mc.UserClaims)
		c.Set("id", mc.UserClaims.ID)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
