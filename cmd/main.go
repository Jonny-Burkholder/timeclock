package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Now serving on port 8080")

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static/")))

	http.Handle("/static/", static)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/home", handleHome)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	//check for login cookie
	//redirect to login or home
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	//look for login cookie
	//redirect if necessary
	//render template
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	//check for cookie... if this were a serious project I'd add it to middleware
	//load
}
