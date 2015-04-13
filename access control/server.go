package server

import (
	"html/template"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles(
	"site/includes/header.html",
	"site/includes/footer.html",
	"site/index.html"))

func init() {
	http.HandleFunc("/", root)
	// http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// func sign(w http.ResponseWriter, r *http.Request) {
// 	err := tmpl.ExecuteTemplate(w, "guestbook.html", r.FormValue("content"))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// }



/*
NOTES:
http://golang.org/pkg/html/template/#Must
https://golang.org/doc/effective_go.html#data
*/