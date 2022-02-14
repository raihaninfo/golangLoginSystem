package main

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/raihaninfo/golangLoginSystem/views"
)

var (
	homeView       *views.View
	aboutView      *views.View
	notFountView   *views.View
	loginView      *views.View
	signupView     *views.View
	forgotPassView *views.View
	fotgotAuthView *views.View
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

func forgotPass(w http.ResponseWriter, r *http.Request) {
	err := forgotPassView.Template.Execute(w, nil)
	FetchError(err)
}

func forgotPassAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	email := r.FormValue("forgotEmail")
	userEmail, err := GetEmail(email)
	FetchError(err)
	if len(userEmail) > 0 {
		isEmail := userEmail[0]["Email"].(string)
		if email == isEmail {
			err := fotgotAuthView.Template.Execute(w, "Check Your Email")
			FetchError(err)

		}
	} else {
		err := fotgotAuthView.Template.Execute(w, "Your account does not exist")
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

func signup(w http.ResponseWriter, r *http.Request) {

	err := signupView.Template.Execute(w, nil)
	FetchError(err)
}

func signupAuth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("fullName")
	email := r.FormValue("email")
	userName := r.FormValue("userName")
	mobile := r.FormValue("mobileNumber")
	password := r.FormValue("Password")

	SignupUser(name, email, userName, mobile, password)
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
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FetchError(err)
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FetchError(err)
}
