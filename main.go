package main

import (
	"coco"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	co := coco.New()
	co.Use(coco.Logger())
	co.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	co.LoadHTMLGlob("templates/*")
	co.Static("/assets", "./static")
	stu1 := &student{Name: "coco", Age: 1}
	stu2 := &student{Name: "Jack", Age: 2}
	co.GET("/", func(c *coco.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	co.GET("/students", func(c *coco.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", coco.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	co.GET("/date", func(c *coco.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", coco.H{
			"title": "gee",
			"now":   time.Date(2019, 8, 17, 0, 0, 0, 0, time.UTC),
		})
	})
	_ = co.Run(":9999")
}
