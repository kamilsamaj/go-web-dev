package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      // DB with users
var dbSessions = map[string]string{} // Session's UUID to Username mapping

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	fmt.Println("listening on http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))

}

func foo(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	var u user
	if userName, ok := dbSessions[c.Value]; ok {
		u = dbUsers[userName]
	}

	// process form submission
	if r.Method == http.MethodPost {
		userName := r.FormValue("username")
		first := r.FormValue("firstname")
		last := r.FormValue("lastname")
		u = user{
			UserName: userName,
			First:    first,
			Last:     last,
		}
		dbSessions[c.Value] = userName
		dbUsers[userName] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}
func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	username, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[username]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)

}
