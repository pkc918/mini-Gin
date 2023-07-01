package main

import (
	"coco"
	"fmt"
	"net/http"
)

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
	_ = co.Run(":9999")
}
