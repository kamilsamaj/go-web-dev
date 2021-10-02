package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseGlob("*.gohtml"))
	data := []string{"hey", "there"}
	fmt.Println(tpl.Execute(os.Stdout, data))

}
