package main

import (
	"net/http"
	"text/template"

	"github.com/Basileus1990/CrowdCompute.git/taskController"
)

// API options
const (
// SHARE = "/share"
)

// returns configured mutex
func setRouting() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", MainPage)
	mux.HandleFunc("/views/static-files/", StaticFiles)
	mux.HandleFunc("/add-task", taskController.AddTask)
	mux.HandleFunc("/get-tasks", taskController.GetTasks)

	return mux
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./views/HTMLTemplates/mainPage.html")
	if err != nil {
		panic(err)
	}
	err = templ.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func StaticFiles(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}
