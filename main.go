package main

import (
	"fmt"
	"net/http"
)

func main() {
	c := coco.New()
	c.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello")
	})
	_ = c.Run(":9999")
}
