package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/sync/singleflight"
	"seat-service/initialization"
	"time"
)

type JWT struct {
	signingKey []byte
}

type UserClaims struct {
	// 可根据需要自行添加字段
	Username string `json:"username"`
	ID       uint   `json:"id"`
}

type CustomClaims struct {
	Buffer int
	UserClaims
	jwt.RegisteredClaims // 内嵌标准的声明
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(initialization.Config.Jwt.SigningKey),
	}
}

// GenToken 生成token
func (j *JWT) GenToken(userClaims UserClaims) (string, error) {
	claims := CustomClaims{
		initialization.Config.Jwt.Buffer,
		userClaims,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(initialization.Config.Jwt.Expires))),
			Issuer:    initialization.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

// ParseToken 解析JWT
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	// 对token对象中的Claim进行类型断言
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims UserClaims) (string, error) {
	Control := &singleflight.Group{}
	v, err, _ := Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.GenToken(claims)
	})
	return v.(string), err
}
