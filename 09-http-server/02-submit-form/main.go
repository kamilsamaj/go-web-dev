package main

import (
	"log"
	"net/http"
	"text/template"
)

type postHandler int

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func (h postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tpl.Execute(w, r.Form)
}

func main() {
	var h postHandler
	http.ListenAndServe(":8080", h)
}
