package main

import (
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{"kamil", 35}
	tpl, _ := template.New("test").Parse("Hello {{ .Name }}, you are {{ .Age }} years old")
	err := tpl.Execute(os.Stdout, p)
	if err != nil {
		panic(err)
	}
}
