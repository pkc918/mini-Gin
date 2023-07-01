package coco

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

// Handle 执行对应路由下的函数
func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "_" + c.Path
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}

// newRouter 生成一个路由表
func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

/*
/ => []
/login/a => [login a]
/login/:router => [login :router]
/login/* => [login *]
/login/*router/a/b => [login *router]
*/
func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			// 遇到 * 后面就不用管了
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

// addRouter 将路由添加到路由表里，构建路由树
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "_" + pattern
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	log.Printf("Route %4s - %s", method, pattern)
	r.handlers[key] = handler // 注册路由表
}

// getRouter 这里就是用户请求进来的时候，遍历树
func (r *router) getRouter(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n != nil {
		// 这个节点下的匹配路径
		parts := parsePattern(n.pattern)
		for i, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[i]
			}
			// *xxx
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[i:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}
