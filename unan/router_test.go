package unan

import (
	"fmt"
	"testing"
)

func fackHandleunc(context *Context) {

}

func buildRoute() *router {
	r := newRoutes()
	r.addRoute("GET", "/users/:id/info", fackHandleunc)
	r.addRoute("PUT", "/users/:id", fackHandleunc)
	r.addRoute("POST", "/users/create", fackHandleunc)
	r.addRoute("DELETE", "/users/:id", fackHandleunc)
	return r
}
func Test_router_addRoute(t *testing.T) {

}

func Test_router_getRouter(t *testing.T) {
	r := buildRoute()
	//n, p := r.getRouter("GET", "/users")
	//t.Log(n, p)

	n, m := r.getRouter("DELETE", "/users/12")
	fmt.Printf("%v, %v", n, m)
	//nodes := make([]*node, 0)
	//assert.Equal(t, len(nodes), 0)
}
