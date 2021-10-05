package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./templates/index.gohtml"))
}

func main() {
	data := "hey there"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, data)
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
