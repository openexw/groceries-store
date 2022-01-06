package unan

import (
	"net/http"
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
	c := newContext(w, req)
	e.router.handle(c)
}
