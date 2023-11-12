package utils

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
	"strings"
)

type CasbinService struct {
	Enforcer *casbin.Enforcer
	Adapter  *gormadapter.Adapter
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
func InitCasbinGorm(db *gorm.DB) (*CasbinService, error) {
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
	return &CasbinService{Adapter: a, Enforcer: enforcer}, nil
}
