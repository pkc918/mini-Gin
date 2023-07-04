package coco

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H {string: interface{}}
type H map[string]interface{}

// Context is the most important part of coco.
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request

	/* request */
	Path   string
	Method string
	Params map[string]string

	/* response */
	StatusCode int

	/* middleware */
	handlers []HandlerFunc
	index    int
}

// newContext 生成一个 Context
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		index:  -1,
	}
}

func (c *Context) Next() {
	c.index++ // 记录当前中间件执行到哪个了,第一次++就是初始化从-1 变为 0
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}

func (c *Context) Fail(code int, err string) {
	c.index = len(c.handlers)
	c.JSON(code, H{"message": err})
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// JSON 序列化返回值
func (c *Context) JSON(code int, res ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	// 将数据 json 序列化
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(res); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// String 序列化返回值
func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// HTML 序列化返回值
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, _ = c.Writer.Write([]byte(html))
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	// 设置响应头
	c.Writer.Header().Set(key, value)
}

// Status 设置响应状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// PostForm 获取请求数据
func (c *Context) PostForm(key string) (value string) {
	value = c.Req.FormValue(key)
	return value
}

// Query 获取 url 上的数据
func (c *Context) Query(key string) (value string) {
	value = c.Req.URL.Query().Get(key)
	return value
}
