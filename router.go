package main

import "net/http"

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	_, ok := session.Values["session"]
	if ok {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	er := loginView.Template.Execute(w, nil)
	FetchError(er)
}

func loginAuth(w http.ResponseWriter, r *http.Request) {

}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Template.Execute(w, nil)
	FetchError(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := aboutView.Template.Execute(w, nil)
	FetchError(err)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}
