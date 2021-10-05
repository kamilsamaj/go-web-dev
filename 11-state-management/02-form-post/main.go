package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	v := r.FormValue("q")
	io.WriteString(w, `
	<html>
	<form method="post">
		<input type="text" name="q">
		<input type="submit">
	</form>
	<br>`+v+`
</html>
`)
}
