package unan

import (
	"net/http"
)

type Engine struct {
	*router
}

type HandlerFunc func(*Context)

func New() *Engine {
	return &Engine{router: newRoutes()}
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
