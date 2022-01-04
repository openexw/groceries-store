package unan

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	//g := New()
	////g.ServeHTTP()
	//http.ListenAndServe(":9091", g)

	g := New()
	g.GET("/", func(c *Context) {
		c.String(http.StatusOK, "ok")
	})

	g.GET("/user/:id", func(c *Context) {
		//c.JSON(http.StatusOK, H{
		//	"code": 200,
		//	"msg":  "ok",
		//	"data": struct{}{},
		//})
		c.JSON(http.StatusOK, H{
			"id": c.Params["id"],
		})
	})

	g.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	g.GET("/login", func(c *Context) {
		c.HTML(http.StatusOK, []byte("<h1>这是登录页面</h1>"))
	})
	t.Fatal(g.Run(":9091"))

}
