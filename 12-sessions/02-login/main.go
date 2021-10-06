package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Username     string
	Fname        string
	Lname        string
	passwordHash []byte
}

var tpl *template.Template
var usersDB = map[string]user{}     // map user's `Username` to a `user` struct
var sessionDB = map[string]string{} // map session's UUID to a Username

const sessionCookieName = "sessionid"

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/signup/", signupHandler)
	http.HandleFunc("/logout/", logoutHandler)
	fmt.Println("listening on http://localhost:8080")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionCookieName)
	if err == http.ErrNoCookie {
		fmt.Println("user not logged in")
	} else if err != nil {
		log.Fatalln("error reading the cookie", sessionCookieName, err)
	}

	var u user
	if c != nil && c.Value != "" {
		fmt.Println("cookie", sessionCookieName, c)
		username := sessionDB[c.Value]
		u = usersDB[username]
	}
	err = tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// resubmission through POST
	if r.Method == http.MethodPost {
		username := r.FormValue("Username")
		passwd := r.FormValue("Password")
		passwdHash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}
		if u, ok := usersDB[username]; ok {
			if err := bcrypt.CompareHashAndPassword(passwdHash, []byte(passwd)); err == nil {
				setCookie(u, w)
				fmt.Println("login successful")
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			} else {
				err := tpl.ExecuteTemplate(w, "login.gohtml", "Incorrect password")
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
		} else if !ok {
			err := tpl.ExecuteTemplate(w, "login.gohtml",
				"User does not exist. Go to /signup/ to create a new account")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	// GET handler
	err := tpl.ExecuteTemplate(w, "login.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		username := r.FormValue("Username")

		// check if a user already exists
		if _, ok := usersDB[username]; ok {
			fmt.Printf("Username '%s' already exists\n", username)
			tpl.ExecuteTemplate(w, "signup.gohtml", "Passwords don't match")
			return
		}

		// check if both passwords match
		passwd := r.FormValue("Password")
		passwd2 := r.FormValue("Password2")
		if passwd != passwd2 {
			fmt.Printf("Password '%s' doesnt match password2 '%s'\n", passwd, passwd2)
			tpl.ExecuteTemplate(w, "signup.gohtml", "Passwords don't match")
			return
		}

		// checks passed, register user
		passwdHash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalln(err)
		}

		u := user{
			Username:     username,
			Fname:        r.FormValue("Firstname"),
			Lname:        r.FormValue("Lastname"),
			passwordHash: passwdHash,
		}
		usersDB[username] = u
		setCookie(u, w)
		fmt.Println("sign-up successful")

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie(sessionCookieName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	delete(sessionDB, c.Value) // remove the cookie from a session db
	http.SetCookie(w, &http.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
	err = tpl.ExecuteTemplate(w, "logout.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setCookie(u user, w http.ResponseWriter) string {
	_uuid := uuid.NewV4().String()
	c := &http.Cookie{
		Name:   sessionCookieName,
		Value:  _uuid,
		MaxAge: 60, // in seconds since now
		Path:   "/",
	}
	sessionDB[_uuid] = u.Username
	http.SetCookie(w, c)
	return _uuid
}
