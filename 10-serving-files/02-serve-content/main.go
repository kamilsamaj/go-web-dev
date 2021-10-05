package main

import (
	"fmt"
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

	http.HandleFunc("/puppy.jpg", servePuppyPic)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func servePuppyPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("./img/puppy.jpg")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.NotFound(w, r)
	}
	http.ServeContent(w, r, fi.Name(), fi.ModTime(), f)
}
