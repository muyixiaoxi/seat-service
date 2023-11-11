package router

// DynamicRoutesResponse1 用于返回路由
type DynamicRoutesResponse1 struct {
	/**
	路由：/system 只需要一层
	例如/system/user 只写user
	*/
	Path string `json:"path"`

	/**
	当前页面的主页地址，一级目录默认 Layout，二级及以后格式： /system/user/index
	*/
	Component string `json:"component"`

	/**
	与Path相同
	*/
	Name string `json:"name"`

	/**
	具体内容
	*/
	Meta RouteContentResponse `json:"meta"`

	/**
	二级路由
	*/
	Children []DynamicRoutesResponse1 `json:"children"`
}

type RouteContentResponse struct {
	/**
	true是隐藏路由，false是不隐藏
	*/
	Hidden bool `json:"hidden"`
	/**
	和路由一样，不用加/  例如 system
	*/
	Icon string `json:"icon"`
	/**
	角色，例如 student 也可以是 1 2 3 4区分
	*/
	Role string `json:"role"`
	/**
	内容，例如 系统管理
	*/
	Title string `json:"title"`
}
