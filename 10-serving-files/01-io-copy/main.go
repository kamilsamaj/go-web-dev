package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/puppy/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg"/></html>`)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><img src="/puppy.jpg"/></html>`)
	})

	http.HandleFunc("/puppy.jpg", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("./img/puppy.jpg")
		defer f.Close()
		if err != nil {
			log.Println(err)
		}
		io.Copy(w, f)
	})
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
