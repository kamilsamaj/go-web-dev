package main

import (
	"fmt"
	"net/http"
)

const (
	cookieName  = "my-cookie"
	cookieValue = "some value"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: cookieValue,
	})
	fmt.Fprintf(w, "Cookie called '%s' has been set to '%s'", cookieName, cookieValue)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("bullshit")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Fprintln(w, "Your cookie:", c)
}
