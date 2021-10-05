package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at foo: %s\n", r.Method)
	fmt.Fprintf(w, "Your request method at foo: %s\n", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at bar: %s\n", r.Method)
	http.Redirect(w, r, "/", 307)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Your request method at barred: %s\n", r.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
