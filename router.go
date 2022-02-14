package main

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if !ok {
		err := loginView.Template.Execute(w, nil)
		FetchError(err)
	}

}

func loginAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := GetUser(email, password)
	FetchError(err)
	if len(user) > 0 {
		session, err := store.Get(r, "login-session")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		session.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 30,
			HttpOnly: true,
		}
		session.Values["username"] = "username"
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		err := loginView.Template.Execute(w, "Please give me right username or password")
		FetchError(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := homeView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func about(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["username"]
	if ok {
		err := aboutView.Template.Execute(w, nil)
		FetchError(err)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}
