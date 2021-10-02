package main

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

func fdateMDY(t time.Time) string {
	return t.Format("Jan-1 2006")
}

var funcMap = template.FuncMap{
	"fdateMDY": fdateMDY,
}

func main() {
	tpl := template.Must(template.New("").Funcs(funcMap).ParseFiles("tpl.gohtml"))
	fmt.Println(tpl.ExecuteTemplate(os.Stdin, "tpl.gohtml", time.Now()))
}
