package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func main() {
	data := struct {
		Wisdom []sage
	}{
		Wisdom: []sage{
			{Name: "Buddha", Motto: "The belief of no beliefs"},
			{Name: "Gandhi", Motto: "Be the change"},
		},
	}
	tpl := template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))

	fmt.Println(tpl.ExecuteTemplate(os.Stdin, "tpl.gohtml", data))
}
