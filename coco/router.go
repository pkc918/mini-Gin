package coco

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

// addRouter 将路由添加到路由表里
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "_" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	r.handlers[key] = handler // 注册路由表
}

// Handle 执行对应路由下的函数
func (r *router) Handle(c *Context) {
	key := c.Method + "_" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// GenerateRouter 生成一个路由表
func GenerateRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}
