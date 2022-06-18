package tinyGin

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type H map[string]interface{}

type Context struct {
    // origin objects
    Writer http.ResponseWriter
    Req *http.Request
    // request info
    Path string
    Method string
    // response info
    StatusCode int
}
