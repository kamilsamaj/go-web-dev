package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-Custom-Header", "This is a value of a custom header")
	fmt.Fprintf(w, "<h1>This is a custom HTTP response</h1>")
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
