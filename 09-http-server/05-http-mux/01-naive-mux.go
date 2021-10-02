package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog/":
		fmt.Fprintf(w, "Wow wow!")
	case "/cat/":
		fmt.Fprintf(w, "Miau Miau!")
	default:
		fmt.Fprintf(w, "I don't know this animal")
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
