package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/ called")
	})
	http.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/dog called")
	})
	http.HandleFunc("/me/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/me/ called")
	})

	fmt.Println("Starting server on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
