package main

import (
	"net/http"
	"tinyGin"
)

/*
(1)
$ curl -i http://localhost:9999/
HTTP/1.1 200 OK
Date: Mon, 12 Aug 2019 16:52:52 GMT
Content-Length: 18
Content-Type: text/html; charset=utf-8
<h1>Hello tiny</h1>
(2)
$ curl "http://localhost:9999/hello?name=YuanHao"
hello YuanHao, you're at /hello
(3)
$ curl "http://localhost:9999/login" -X POST -d 'username=YuanHao&password=123456'
{"password":"1234","username":"YuanHao"}
(4)
$ curl "http://localhost:9999/xxx"
404 NOT FOUND: /xxx
*/

func main() {
	r := tinyGin.New()

	r.GET("/", func(c *tinyGin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello TinyGin</h1>")
	})

	r.GET("/hello", func(c *tinyGin.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *tinyGin.Context) {
		c.JSON(http.StatusOK, tinyGin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
