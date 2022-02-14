package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"github.com/raihaninfo/golangLoginSystem/views"
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
	signupView = views.NewView("views/fron-end/signin.gohtml")
	forgotPassView= views.NewView("views/fron-end/forgotpass.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)
	r.HandleFunc("/signup", signup)
	r.HandleFunc("/signupauth", signupAuth)
	r.HandleFunc("/logout", logout)
	r.HandleFunc("/forgot_pass", forgotPass)

	fmt.Println("Listening port :8082")
	http.ListenAndServe(":8082", r)
}

func FetchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
