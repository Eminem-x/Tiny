package main

/*
(1) index
curl -i http://localhost:9999/index
HTTP/1.1 200 OK
Date: Sun, 01 Sep 2019 08:12:23 GMT
Content-Length: 19
Content-Type: text/html; charset=utf-8
<h1>Index Page</h1>

(2) v1
$ curl -i http://localhost:9999/v1/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 18:11:07 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello tinyGin</h1>

(3)
$ curl "http://localhost:9999/v1/hello?name=YuanHao"
hello YuanHao, you're at /v1/hello

(4)
$ curl "http://localhost:9999/v2/hello/YuanHao"
hello YuanHao, you're at /hello/YuanHao

(5)
$ curl "http://localhost:9999/v2/login" -X POST -d 'username=YuanHao&password=1234'
{"password":"1234","username":"YuanHao"}

(6)
$ curl "http://localhost:9999/hello"
404 NOT FOUND: /hello
*/

import (
	"net/http"

	"tinyGin"
)

func main() {
	r := tinyGin.New()
	r.GET("/index", func(c *tinyGin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *tinyGin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello tinyGin</h1>")
		})

		v1.GET("/hello", func(c *tinyGin.Context) {
			// expect /hello?name=YuanHao
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *tinyGin.Context) {
			// expect /hello/YuanHao
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *tinyGin.Context) {
			c.JSON(http.StatusOK, tinyGin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
