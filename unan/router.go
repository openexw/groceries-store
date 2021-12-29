package unan

import (
	"fmt"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRoutes() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
func (r *router) addRoute(method, path string, handler HandlerFunc) {
	key := method + ":" + path
	r.handlers[key] = handler
}
func (r *router) GET(path string, handler HandlerFunc) {
	r.addRoute(http.MethodGet, path, handler)
}

func (r *router) POST(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPost, path, handler)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//url := req.URL.Path
	//switch url {
	//case "/":
	//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	//case "/ping":
	//	for k, header := range req.Header {
	//		fmt.Fprintf(w, "Header[%q] = %q\n", k, header)
	//	}
	//default:
	//	fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	//}
	key := req.Method + ":" + req.URL.Path
	if handler, ok := r.handlers[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
