package main

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello tinyGin</h1>

(2)
$ curl "http://localhost:9999/hello?name=YuanHao"
hello YuanHao, you're at /hello

(3)
$ curl "http://localhost:9999/hello/YuanHao"
hello YuanHao, you're at /hello/YuanHao

(4)
$ curl "http://localhost:9999/assets/css/YuanHao.css"
{"filepath":"css/YuanHao.css"}

(5)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

import (
	"net/http"
	"tinyGin"
)

func main() {
	r := tinyGin.New()
	r.GET("/", func(c *tinyGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello tinyGin</h1>")
	})

	r.GET("/hello", func(c *tinyGin.Context) {
		// expect /hello?name=YuanHao
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *tinyGin.Context) {
		// expect /hello/YuanHao
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *tinyGin.Context) {
		c.JSON(http.StatusOK, tinyGin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
