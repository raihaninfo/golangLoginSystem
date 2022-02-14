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
	// jj, err := GetUser("raihanmahmudi35@gmail.com", "raihan123")
	// FetchError(err)
	// jjs := jj[0]["Email"]
	// fmt.Println(jjs)

	homeView = views.NewView("views/fron-end/index.gohtml")
	aboutView = views.NewView("views/fron-end/about.gohtml")
	notFountView = views.NewView("views/fron-end/notfount.gohtml")
	loginView = views.NewView("views/fron-end/login.gohtml")
	r := mux.NewRouter()
	r.PathPrefix("/asset/").Handler(http.StripPrefix("/asset/", http.FileServer(http.Dir("views/fron-end/asset"))))
	r.NotFoundHandler = http.HandlerFunc(notFount)
	r.HandleFunc("/login", login)
	r.HandleFunc("/loginauth", loginAuth)
	r.HandleFunc("/", home)
	r.HandleFunc("/about", about)

	fmt.Println("Listening port :8082")
	http.ListenAndServe(":8082", r)
}

func FetchError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
