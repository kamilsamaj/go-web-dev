package main

import (
	"fmt"
	"net/http"
)

fmt.Println("listening on http://localhost:8080")

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets/")))

	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
