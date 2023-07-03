package coco

import (
	"net/http"
)

// HandlerFunc  defines the request handler type by coco
type HandlerFunc func(c *Context)

// RouterGroup 路由分组
type RouterGroup struct {
	prefix      string        // 前缀：作用在哪个路由
	middlewares []HandlerFunc // 中间件处理函数
	parent      *RouterGroup
	engine      *Engine
}

// Engine implement the interface of ServeHTTP
type Engine struct {
	*RouterGroup
	groups []*RouterGroup
	router *router
}

// New return a constructor of coco.Engine
func New() *Engine {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group 分组
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

// ServeHTTP 拦截所有的http请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// Run defines the method to start a http server
func (engine *Engine) Run(address string) (err error) {
	// 所有的请求都会走 engine
	return http.ListenAndServe(address, engine)
}

// GET 默认一个 get 请求
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

// POST 默认一个 post 请求
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRoute(method, pattern, handler)
}
