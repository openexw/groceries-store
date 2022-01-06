package unan

import "net/http"

type RouteGroup struct {
	prefix      string        // 分组前缀
	middlewares []HandlerFunc // 当前分组支持的中间件列表
	parent      *RouteGroup   // 支持分组嵌套
	engine      *Engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
	e := group.engine

	g := &RouteGroup{
		prefix:      group.prefix + prefix,
		middlewares: nil,
		parent:      group,
		engine:      e,
	}

	e.groups = append(e.groups, g)
	return g
}

func (group *RouteGroup) addRoute(methods, comp string, handler HandlerFunc) {
	path := group.prefix + comp
	group.engine.router.addRoute(methods, path, handler)
}

func (group *RouteGroup) GET(path string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, path, handler)
}

func (group *RouteGroup) POST(path string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, path, handler)
}

func (group *RouteGroup) DELETE(path string, handler HandlerFunc) {
	group.addRoute(http.MethodDelete, path, handler)
}

func (group *RouteGroup) PUT(path string, handler HandlerFunc) {
	group.addRoute(http.MethodPut, path, handler)
}

func (group *RouteGroup) PATCH(path string, handler HandlerFunc) {
	group.addRoute(http.MethodPatch, path, handler)
}

func (group *RouteGroup) OPTIONS(path string, handler HandlerFunc) {
	group.addRoute(http.MethodOptions, path, handler)
}

func (group *RouteGroup) HEAD(path string, handler HandlerFunc) {
	group.addRoute(http.MethodHead, path, handler)
}
