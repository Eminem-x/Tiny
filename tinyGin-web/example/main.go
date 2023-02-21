package main

import (
	"tinyGin"
)

func main() {
	r := tinyGin.Default()
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
