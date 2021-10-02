package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	// package http has a default DefaultServeMux
	http.HandleFunc("/hi/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi %s!", strings.TrimPrefix(r.URL.Path, "/hi/"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the homepage")
	})

	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
