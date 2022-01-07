package unan

import (
	"net/http"
	"strings"
)

type Engine struct {
	*RouteGroup               // 嵌套，Engine 拥有 RouteGroup 的所有能力
	router      *router       // 路由
	groups      []*RouteGroup // 存储所有的分组
}

type HandlerFunc func(*Context)

func New() *Engine {
	e := &Engine{
		router: newRoutes(),
	}
	e.RouteGroup = &RouteGroup{engine: e}
	e.groups = []*RouteGroup{e.RouteGroup}
	return e
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//key := req.Method + ":" + req.URL.Path
	//if handler, ok := r.handlers[key]; ok {
	//	handler(w, req)
	//} else {
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		// 根据 URL 搜索前缀判断是否是同一个 group
		// 将该 group 下的中间件放进 context 中
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares

	e.router.handle(c)
}
