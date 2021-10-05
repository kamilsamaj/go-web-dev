package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey there!")
}

func main() {
	http.Handle("/", http.HandlerFunc(handler))

	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
