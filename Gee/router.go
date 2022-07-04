package Gee

import (
	"net/http"
	"strings"
)

type router struct {
	handlers map[string]HandleFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

func (r *router) addRoute(method string,pattern string,handler HandleFunc){
	var key strings.Builder
	key.WriteString(method)
	key.WriteString("-")
	key.WriteString(pattern)
	r.handlers[key.String()] = handler
}

func (r *router) handle(c *Context) {
	var key strings.Builder
	key.WriteString(c.R.Method)
	key.WriteString("-")
	key.WriteString(c.R.URL.Path)
	if handler, ok := r.handlers[key.String()]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

