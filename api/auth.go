package api

import (
	"github.com/gin-gonic/gin"
	"seat-service/response"
	"seat-service/utils"
)

type AuthApi struct {
}

var casbinMethod utils.CasbinMethod

// GetUsers 获取所有用户
func (a *AuthApi) GetUsers(context *gin.Context) {
	resp.Success(context, response.CodeSuccess, casbinMethod.GetUsers())
}

// GetRoles 获取所有角色组
func (a *AuthApi) GetRoles(context *gin.Context) {
	resp.Success(context, response.CodeSuccess, casbinMethod.GetRoles())
}

// GetRolePolicy 获取所有角色组的策略
func (a *AuthApi) GetRolePolicy(context *gin.Context) {
	roles, err := casbinMethod.GetRolePolicy()
	if err != nil {

	} else {
		resp.Success(context, response.CodeSuccess, roles)
	}
}

/* 创建角色策略
  type RolePolicy struct {
    RoleName string `gorm:"column:v0"`
    Url      string `gorm:"column:v1"`
    Method   string `gorm:"column:v2"`
}
*/

func (a *AuthApi) CreateRolePolicy(context *gin.Context) {
	var p utils.RolePolicy
	err := context.ShouldBindJSON(&p)
	if err != nil {
		return
	}
	err = casbinMethod.CreateRolePolicy(p)
	if err != nil {

	} else {
		resp.Success(context, response.CodeSuccess, nil)
	}
}

/* 删除角色组策略
  type RolePolicy struct {
    RoleName string `gorm:"column:v0"`
    Url      string `gorm:"column:v1"`
    Method   string `gorm:"column:v2"`
}
*/

func (a *AuthApi) DeleteRolePolicy(context *gin.Context) {

	var p utils.RolePolicy
	err := context.ShouldBindJSON(&p)
	if err != nil {
		return
	}
	err = casbinMethod.DeleteRolePolicy(p)
	if err != nil {

	} else {

	}
}

// UpdateUserRole 添加用户组策略, /casbin/user-role?username=leo&rolename=admin
func (a *AuthApi) UpdateUserRole(context *gin.Context) {

	username := context.Query("username")
	rolename := context.Query("rolename")
	err := casbinMethod.UpdateUserRole(username, rolename)
	if err != nil {

	} else {

	}
}

// DeleteUserRole 从组中删除用户, /casbin/user-role?username=leo&rolename=admin
func (a *AuthApi) DeleteUserRole(context *gin.Context) {
	username := context.Query("username")
	rolename := context.Query("rolename")
	err := casbinMethod.DeleteUserRole(username, rolename)
	if err != nil {
		context.String(500, "从组中删除用户失败: "+err.Error())
	} else {
		context.JSON(200, "成功!")
	}
}
