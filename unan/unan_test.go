package unan

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	host = "http://127.0.0.1"
	port = ":9091"
)

func TestNew(t *testing.T) {
	setupUnanServe(t)
	//g := New()
	////g.ServeHTTP()
	//http.ListenAndServe(":9091", g)
	//
	//g := New()
	//g.GET("/", func(c *Context) {
	//	c.String(http.StatusOK, "ok")
	//})
	//
	//g.GET("/user/:id", func(c *Context) {
	//	//c.JSON(http.StatusOK, H{
	//	//	"code": 200,
	//	//	"msg":  "ok",
	//	//	"data": struct{}{},
	//	//})
	//	c.JSON(http.StatusOK, H{
	//		"id": c.Params["id"],
	//	})
	//})
	//
	//g.GET("/ping", func(c *Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	//
	//g.GET("/login", func(c *Context) {
	//	c.HTML(http.StatusOK, []byte("<h1>这是登录页面</h1>"))
	//})
	//t.Fatal(g.Run(":9091"))

}

func onlyForUsers() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		log.Printf("[%d] %s in %v for group users", c.statusCode, c.req.RequestURI, time.Since(t))
	}
}

func setupUnanServe(t *testing.T) {
	g := New()
	g.Use(Logger(), Recovery()) // 使用中间件
	g.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	users := g.Group("/users")
	users.Use(onlyForUsers())
	users.GET("/:id", func(c *Context) {
		c.String(http.StatusOK, fmt.Sprintf("info: users/%s", c.Param("id")))
	})

	users.POST("/create", func(c *Context) {
		c.String(http.StatusOK, "created")
	})

	users.DELETE("/del", func(c *Context) {
		c.String(http.StatusOK, "deleted")
	})

	users.PUT("/:id", func(c *Context) {
		c.String(http.StatusOK, fmt.Sprintf("modify: users/%s", c.Param("id")))
	})
	t.Fatal(g.Run(port))
}

func getReqURL(uri string) string {
	return host + port + uri
}

func TestRouteGroup_Group(t *testing.T) {

	got := httptest.NewRequest(http.MethodGet, getReqURL("/users/1"), nil)
	body, err := io.ReadAll(got.Body)
	if err != nil {
		t.Errorf("ReadAll: %v", err)
	}
	assert.Equal(t, "info: users/1", string(body))

}
