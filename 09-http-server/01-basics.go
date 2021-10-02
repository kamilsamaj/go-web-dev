package main

import (
	"fmt"
	"net/http"
)

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "This got written by my custom HTTP handler\n")
}

func main() {
	var h myHandler
	http.ListenAndServe(":8080", h)
}
