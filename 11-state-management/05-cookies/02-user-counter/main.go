package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const cookieName = "__user_counter"

func main() {
	http.HandleFunc("/", incrementCookie)
	fmt.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func incrementCookie(w http.ResponseWriter, r *http.Request) {
	// first, try to read the cookie
	c, err := r.Cookie("__user_counter")
	if err != nil {
		if strings.Contains(err.Error(), "cookie not present") {
			http.SetCookie(w, &http.Cookie{
				Name:  cookieName,
				Value: "1",
			})
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// if we didn't return, increment the cookie value
	counterVal, err := strconv.Atoi(c.Value)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	newCookie := http.Cookie{
		Name:  cookieName,
		Value: strconv.Itoa(counterVal + 1),
	}

	http.SetCookie(w, &newCookie)

	fmt.Fprintf(w, "Cookie %s set to %d", cookieName, counterVal+1)
}
