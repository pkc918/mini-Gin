package coco

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		begin := time.Now()
		c.Next()
		elapse := time.Since(begin)
		log.Printf("%s use time %d ms", c.Path, elapse)
	}
}
