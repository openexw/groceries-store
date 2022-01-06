package unan

import (
	"net/http"
	"strings"
)

type router struct {
	// handlers 存放 route 的处理 handler，key 为 method:uri，如：GET:/user/create
	handlers map[string]HandlerFunc
	roots    map[string]*node
}

func newRoutes() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots:    make(map[string]*node),
	}
}
func (r *router) addRoute(method, path string, handler HandlerFunc) {
	// add node
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = newNode()
	}
	r.roots[method].insertChild(path)

	// add handler
	key := method + ":" + path
	r.handlers[key] = handler
}

func (r *router) getRouter(method, path string) (*node, map[string]string) {
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(path)
	if n == nil {
		return nil, nil
	}

	if n.wildChild {
		reqParts := parsePath(path)
		params := make(map[string]string)
		parts := parsePath(n.fullPath)
		for index, part := range parts {
			// route: /hello/:name
			// req:   /hello/Jhon
			// 将 :name 与 value 进行映射
			if part[0] == ':' {
				name := part[1:]
				params[name] = reqParts[index]
			}

			// 处理 * 这样的
			if part[0] == '*' && len(part) > 1 {
				name := part[1:]
				params[name] = strings.Join(reqParts[index:], "/")
			}
		}
		return n, params
	}
	return n, nil
}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + ":" + n.fullPath
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
