package unan

import (
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

func (r *router) DELETE(path string, handler HandlerFunc) {
	r.addRoute(http.MethodDelete, path, handler)
}

func (r *router) PUT(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPut, path, handler)
}

func (r *router) PATCH(path string, handler HandlerFunc) {
	r.addRoute(http.MethodPatch, path, handler)
}

func (r *router) OPTIONS(path string, handler HandlerFunc) {
	r.addRoute(http.MethodOptions, path, handler)
}

func (r *router) HEAD(path string, handler HandlerFunc) {
	r.addRoute(http.MethodHead, path, handler)
}

func (r *router) handle(c *Context) {
	key := c.Method + ":" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
