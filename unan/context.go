package unan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	w      http.ResponseWriter
	req    *http.Request
	Method string
	Path   string
	//handler HandlerFunc
	Params     map[string]string // 请求参数
	statusCode int

	// 中间件及本身的函数
	handlers []HandlerFunc
	index    int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w: w,
		req:    req,
		Method: req.Method,
		Path:   req.URL.Path,
		index:  -1,
	}
}

// Next 中间件 next
func (c *Context) Next() {
	c.index++
	l := len(c.handlers)
	for ; c.index < l; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) SetStatusCode(statusCode int) {
	c.statusCode = statusCode
	c.w.WriteHeader(statusCode)
}

func (c *Context) SetHeader(key, value string) {
	c.w.Header().Set(key, value)
}

func (c *Context) String(status int, format string, values ...interface{}) {
	c.SetStatusCode(status)
	c.SetHeader("Content-Type", "text/plain")
	c.w.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(status int, data interface{}) {
	c.SetStatusCode(status)
	c.SetHeader("Content-Type", "application/json")

	encoder := json.NewEncoder(c.w)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(c.w, err.Error(), http.StatusInternalServerError)
	}
}
func (c *Context) HTML(status int, html []byte) {
	c.SetStatusCode(status)
	c.SetHeader("Content-Type", "text/html")
	c.w.Write(html)
}

func (c *Context) Param(key string) string {
	val := c.Params[key]
	return val
}
