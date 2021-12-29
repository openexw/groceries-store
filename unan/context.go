package unan

import "net/http"

type Context struct {
	w      http.ResponseWriter
	req    *http.Request
	Method string
	path   string
	//handler HandlerFunc
}
