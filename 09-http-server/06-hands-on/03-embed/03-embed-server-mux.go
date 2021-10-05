package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

var indexTpl *template.Template

//go:embed templates/*
var templates embed.FS

func init() {
	indexTpl = template.Must(template.ParseFS(templates, "templates/index.gohtml"))
}

func main() {
	data := "hey there"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexTpl.Execute(w, data)
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
