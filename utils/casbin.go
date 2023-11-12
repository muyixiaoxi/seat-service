package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"strings"
)

type CasbinMethod struct {
	Enforcer *casbin.Enforcer
	Adapter  *gormadapter.Adapter
}

// RolePolicy (RoleName, Url, Method) 对应于 `CasbinRule` 表中的 (v0, v1, v2)
type RolePolicy struct {
	RoleName string `gorm:"column:v0" json:"role_name"`
	Url      string `gorm:"column:v1" json:"url"`
	Method   string `gorm:"column:v2" json:"method"`
}

type User struct {
	UserName  string
	RoleNames []string
}

// KeyMatchFunc 自定义Casbin匹配规则
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return (bool)(KeyMatch(name1, name2)), nil
}

func KeyMatch(key1, key2 string) bool {
	i := strings.Index(key2, "*")
	if i == -1 {
		return key1 == key2
	}

	if len(key1) > i {
		return key1[:i] == key2[:i]
	}

	return key1 == key2[:i]
}

// InitCasbinGorm 初始化Casbin Gorm适配器
func InitCasbinGorm(db *gorm.DB) (*CasbinMethod, error) {
	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		return nil, err
	}
	enforcer, err := casbin.NewEnforcer("./model.config", a)

	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	//将自定义权限匹配规则加入权限认证器
	enforcer.AddFunction("my_func", KeyMatchFunc)
	return &CasbinMethod{Adapter: a, Enforcer: enforcer}, nil
}

// GetRoles 获取所有角色组
func (c *CasbinMethod) GetRoles() []string {
	return c.Enforcer.GetAllRoles()
}

// GetRolePolicy 获取所有角色组权限
func (c *CasbinMethod) GetRolePolicy() (roles []RolePolicy, err error) {
	err = c.Adapter.GetDb().Model(&gormadapter.CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// CreateRolePolicy 创建角色组权限, 已有的会忽略
func (c *CasbinMethod) CreateRolePolicy(r RolePolicy) error {
	// 不直接操作数据库，利用enforcer简化操作
	err := c.Enforcer.LoadPolicy()
	if err != nil {
		return err
	}
	_, err = c.Enforcer.AddPolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// UpdateRolePolicy 修改角色组权限
func (c *CasbinMethod) UpdateRolePolicy(old, new RolePolicy) error {
	_, err := c.Enforcer.UpdatePolicy([]string{old.RoleName, old.Url, old.Method},
		[]string{new.RoleName, new.Url, new.Method})
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// DeleteRolePolicy 删除角色组权限
func (c *CasbinMethod) DeleteRolePolicy(r RolePolicy) error {
	_, err := c.Enforcer.RemovePolicy(r.RoleName, r.Url, r.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// GetUsers 获取所有用户以及关联的角色
func (c *CasbinMethod) GetUsers() (users []User) {
	p := c.Enforcer.GetGroupingPolicy()
	usernameUser := make(map[string]*User, 0)
	for _, _p := range p {
		username, usergroup := _p[0], _p[1]
		if v, ok := usernameUser[username]; ok {
			usernameUser[username].RoleNames = append(v.RoleNames, usergroup)
		} else {
			usernameUser[username] = &User{UserName: username, RoleNames: []string{usergroup}}
		}
	}
	for _, v := range usernameUser {
		users = append(users, *v)
	}
	return
}

// UpdateUserRole 角色组中添加用户, 没有组默认创建
func (c *CasbinMethod) UpdateUserRole(username, rolename string) error {
	_, err := c.Enforcer.AddGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// DeleteUserRole 角色组中删除用户
func (c *CasbinMethod) DeleteUserRole(username, rolename string) error {
	_, err := c.Enforcer.RemoveGroupingPolicy(username, rolename)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy()
}

// CanAccess 验证用户权限
func (c *CasbinMethod) CanAccess(username, url, method string) (ok bool, err error) {
	return c.Enforcer.Enforce(username, url, method)
}
