package handler

import (
	"fmt"
	"log"
	"net/http"
	"tinyGin"
)

func Ping(c *tinyGin.Context) {
	c.String(http.StatusOK, "pong")
}

func Context(c *tinyGin.Context) {
	for k, v := range c.Req.Header {
		_, err := fmt.Fprintf(c.Writer, "<p>Header[%q] = %q</p>\n", k, v)
		if err != nil {
			c.String(http.StatusInternalServerError, "fmt err: %#v", err)
			c.Next()
		}
	}
}

func Login(c *tinyGin.Context) {
	if c.Query("name") == "ycx" {
		c.String(http.StatusOK, "success")
	} else {
		c.String(http.StatusUnauthorized, "auth fail")
	}
}

func Logger(c *tinyGin.Context) {
	log.Printf("request url: %#v", c.Req.URL)
	c.Next()
}

func Panic(c *tinyGin.Context) {
	panic("panic")
}
