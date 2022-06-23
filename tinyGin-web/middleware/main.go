package main

/*
(1) global middleware Logger
$ curl http://localhost:9999/
<h1>Hello tinyGin</h1>

>>> log
2022/06/23 01:37:38 [200] / in 3.14µs
*/

/*
(2) global + group middleware
$ curl http://localhost:9999/v2/hello/YuanHao
{"message":"Internal Server Error"}

>>> log
2022/06/23 01:38:48 [200] /v2/hello/YuanHao in 61.467µs for group v2
2022/06/23 01:38:48 [200] /v2/hello/YuanHao in 281µs
*/

import (
	"log"
	"net/http"
	"time"

	"tinyGin"
)

func onlyForV2() tinyGin.HandlerFunc {
	return func(c *tinyGin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := tinyGin.New()
	r.Use(tinyGin.Logger()) // global middleware

	r.GET("/", func(c *tinyGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tinyGin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *tinyGin.Context) {
			// expect /hello/YuanHao
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
