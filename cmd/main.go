package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Jonny-Burkholder/timeclock/internal/tools"
)

func main() {
	err := userMap.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("Now serving on port 8080")

	static := http.StripPrefix("/static/", http.FileServer(http.Dir("../web/static/")))

	http.Handle("/static/", static)
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/submit-pin", handleSubmitPin)
	http.HandleFunc("/home/", handleHome)
	http.HandleFunc("/clock-in/", handleClockIn)
	http.HandleFunc("/clock-out/", handleClockOut)
	http.HandleFunc("/shift-report/", handleShiftReport)

	http.ListenAndServe(":8080", nil)
}

var userMap = tools.NewUserMap()

var templates = template.Must(template.New("*").Funcs(funcMap).ParseGlob("../web/templates/*.html"))

var funcMap = template.FuncMap{
	"DisplayTime":  tools.DisplayTime,
	"DisplayShift": tools.DisplayShift,
}

func renderTemplate(w http.ResponseWriter, tmpl string, p tools.AnyPage) {
	buffer := tools.GetBuf()
	err := templates.ExecuteTemplate(buffer, tmpl+".html", p)
	if err != nil {
		buffer.WriteTo(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buffer.WriteTo(w)
	tools.PutBuf(buffer)
	return
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusFound)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	page := &tools.Page{}
	renderTemplate(w, "login", page)
}

func handleSubmitPin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := userMap.CheckPin(r.FormValue("pin_field"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "login", http.StatusFound)
	} else {
		http.Redirect(w, r, "home/"+user.Username, http.StatusFound)
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/home/"):]
	user, err := userMap.LoadUser(username)
	if err != nil {
		http.Error(w, "User Not Found", 503)
	}
	renderTemplate(w, "home", user)
}

func handleShiftReport(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/shift-report/"):]
	user, err := userMap.LoadUser(username)
	if err != nil {
		http.Error(w, "User Not Found", 503)
	}
	renderTemplate(w, "shift-report", user)
}

func handleClockIn(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/clock-in/"):]
	user, err := userMap.LoadUser(username)
	if err != nil {
		http.Error(w, "User Not Found", 503)
		return
	}
	user.StartShift()
	http.Redirect(w, r, "/login", http.StatusFound)
}

func handleClockOut(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Path[len("/clock-out/"):]
	user, err := userMap.LoadUser(username)
	if err != nil {
		http.Error(w, "User Not Found", 503)
		return
	}
	user.EndShift()
	http.Redirect(w, r, "/shift-report/"+user.Username, http.StatusFound)
}
