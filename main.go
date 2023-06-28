package main

import (
	"coco"
	"fmt"
	"net/http"
)

func main() {
	co := coco.New()
	co.POST("/login", func(c *coco.Context) {
		fmt.Println(c.Path)
		c.JSON(http.StatusOK, coco.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	_ = co.Run(":9999")
}
