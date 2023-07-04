package main

import (
	"coco"
	"fmt"
	"log"
	"net/http"
	"time"
)

func middlewareByV2() coco.HandlerFunc {
	return func(c *coco.Context) {
		t := time.Now()
		c.Fail(500, "Internet Server Error")
		//c.Status(200)
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	co := coco.New()
	co.GET("/", func(c *coco.Context) {
		c.HTML(http.StatusOK, "<h1>Hello coco</h1>")
	})

	co.GET("/userInfo", func(c *coco.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	co.POST("/login", func(c *coco.Context) {
		fmt.Println(c.Path)
		c.JSON(http.StatusOK, coco.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	co.GET("/assets/*filepath", func(c *coco.Context) {
		c.JSON(http.StatusOK, coco.H{
			"filepath": c.Param("filepath"),
		})
	})

	v1 := co.Group("/v1")
	{
		v1.GET("/", func(c *coco.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *coco.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := co.Group("/v2")
	v2.Use(middlewareByV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *coco.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	_ = co.Run(":9999")
}
