package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("./resources/dog.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/resources/puppy.jpg", dogImg)

	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	http.NotFoundHandler()
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "This is from dog")
}

func dogImg(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./resources/puppy.jpg")
}
