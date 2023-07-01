package coco

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/api/:name", nil)
	r.addRoute("GET", "/api/user/login", nil)
	r.addRoute("GET", "/api/user/:name", nil)
	r.addRoute("GET", "/api/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	assert.Equal(t, parsePattern("/"), []string{})
	assert.Equal(t, parsePattern("/api/:name"), []string{"api", ":name"})
	assert.Equal(t, parsePattern("/api/user/login"), []string{"api", "user", "login"})
	assert.Equal(t, parsePattern("/api/user/:name"), []string{"api", "user", ":name"})
	assert.Equal(t, parsePattern("/assets/*filePath"), []string{"api", "*filePath"})
}

func TestGetRouter(t *testing.T) {
	path := "/api/user/register"
	r := newTestRouter()
	n, ps := r.getRouter("GET", path)
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}
	if n.pattern != "/api/user/:name" {
		t.Fatal("should match path => /api/user/:name")
	}
	fmt.Printf("matched path: %s, params[name]: %s\n", n.pattern, ps["name"])
}
