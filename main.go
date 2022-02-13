package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raihaninfo/golangLoginSystem/views"
)

var (
	homeView     *views.View
	aboutView    *views.View
	notFountView *views.View
	loginView    *views.View
)

var store = sessions.NewCookieStore([]byte("secret-password"))

func init() {
	dbcon()
}

func main() {
	homeView = views.NewView("views/fron-end/index.gohtml")
	aboutView = views.NewView("views/fron-end/about.gohtml")
	notFountView = views.NewView("views/fron-end/notfount.gohtml")
	loginView = views.NewView("views/fron-end/login.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)

	fmt.Println("Listening port :8082")
	http.ListenAndServe(":8082", r)
}

func login(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "login-session")
	FeatchError(err)
	_, ok := session.Values["session"]
	if ok {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
	er := loginView.Template.Execute(w, nil)
	FeatchError(er)
}

func home(w http.ResponseWriter, r *http.Request) {
	err := homeView.Template.Execute(w, nil)
	FeatchError(err)
}

func about(w http.ResponseWriter, r *http.Request) {
	err := aboutView.Template.Execute(w, nil)
	FeatchError(err)
}

func notFount(w http.ResponseWriter, r *http.Request) {
	err := notFountView.Template.Execute(w, nil)
	FeatchError(err)
}

func FeatchError(err error) {
	if err != nil {
		panic(err)
	}
}
