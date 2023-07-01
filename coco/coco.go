package coco

import (
	"net/http"
)

// HandlerFunc  defines the request handler type by coco
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New return a constructor of coco.Engine
func New() *Engine {
	return &Engine{router: newRouter()}
}

// ServeHTTP 拦截所有的http请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := GenerateContext(w, req)
	engine.router.handle(c)
}

// Run defines the method to start a http server
func (engine *Engine) Run(address string) (err error) {
	// 所有的请求都会走 engine
	return http.ListenAndServe(address, engine)
}

// GET 默认一个 get 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 默认一个 post 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 中转一下
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}
