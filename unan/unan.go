package unan

import (
	"net/http"
)

type Engine struct {
	*router
}

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

func New() *Engine {
	return &Engine{router: newRoutes()}
}

func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}
