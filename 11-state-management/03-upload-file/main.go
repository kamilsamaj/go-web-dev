package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	var s string
	fmt.Printf("received %s method\n", r.Method)

	if r.Method == http.MethodPost {

		f, header, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", header, "\nerr", err)

		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
		cap()
	}
	fmt.Fprintf(w, `
<html>
<form method="post" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
</form>
<br>%s
</html>
`, s)
}
