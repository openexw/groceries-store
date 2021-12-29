package unan

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	//g := New()
	////g.ServeHTTP()
	//http.ListenAndServe(":9091", g)

	g := New()
	g.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	g.POST("/user/create", func(w http.ResponseWriter, req *http.Request) {
		for k, header := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, header)
		}
	})

	g.GET("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "pong")
	})
	g.Run(":9091")
}
