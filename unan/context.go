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

	statusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{w: w, req: req, Method: req.Method, Path: req.URL.Path}
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
