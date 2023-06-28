package coco

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc  defines the request handler type by coco
type HandlerFunc func(c *Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// ServeHTTP 拦截所有的http请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("12321312312")
	c := GenerateContext(w, req)
	// 先获取请求的路径，按照格式拼接成需要的 key
	key := req.Method + "_" + req.URL.Path
	// 从对应的路由表中取方法，取得到就执行，取不到就说明用户访问了一个不存在的路由
	if handler, ok := engine.router[key]; ok {
		handler(c)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// New return a constructor of coco.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// Run defines the method to start a http server
func (engine *Engine) Run(address string) (err error) {
	// 所有的请求都会走 engine
	return http.ListenAndServe(address, engine)
}

// addRouter 将路由添加到路由表里
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "_" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler // 注册路由表
}

// GET 默认一个 get 请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 默认一个 post 请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}
