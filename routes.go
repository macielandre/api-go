package main

import (
	"fmt"
	"net/http"
)

func test(req http.ResponseWriter, r *http.Request) {
	fmt.Println(req, "hello")
}
