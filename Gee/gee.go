package Gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W http.ResponseWriter
	R *http.Request
	Method string
	Path string
	StatusCode int
}

type HandleFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine{
	return &Engine{router: newRouter()}
}

func NewContext(w http.ResponseWriter,r *http.Request) *Context{
	return &Context{
		W:w,
		R:r,
		Method: r.Method,
		Path: r.URL.Path,
	}
}

func (engine *Engine) addRoute(method string,pattern string,handler HandleFunc){
	engine.router.addRoute(method,pattern,handler)
}

func (engine *Engine) GET(pattern string,handler HandleFunc){
	engine.addRoute("GET",pattern,handler)
}

func (engine *Engine) Run(addr string){
	http.ListenAndServe(addr,engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter,r *http.Request){
	c := NewContext(w,r)
	engine.router.handle(c)
}

func (c *Context) PostForm(key string) string{
	return c.R.FormValue(key)
}

func (c *Context) Get(key string) string{
	return c.R.URL.Query().Get(key)
}

func (c *Context) Status(code int){
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) SetHeader(key string,value string){
	c.W.Header().Set(key,value)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.W, err.Error(), 500)
	}
}


