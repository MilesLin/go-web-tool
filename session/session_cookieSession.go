package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
)

var store *sessions.CookieStore

func init() {
	
	store = sessions.NewCookieStore([]byte("secret-key"))
}

func main() {
	http.HandleFunc("/home", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	fmt.Println("session server run on port 8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = false
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "logged out.")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["auth"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "logged in")
}

func home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	auth := session.Values["auth"]
	if auth != nil {
		isAuth, ok := auth.(bool)
		if ok && isAuth {
			fmt.Fprintln(w, "Home Page")
		} else {
			http.Error(w, "unauthorizeed", http.StatusUnauthorized)
			return
		}
	} else {
		http.Error(w, "unauthorizeed", http.StatusUnauthorized)
		return
	}

}
