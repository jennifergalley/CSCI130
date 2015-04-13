package main

import (
  "html/template"
  "net/http"
)

var mytemp *template.Template

type User struct{
    Name, Email, Message string
}

func init() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/formhandler", formhandler)
  http.HandleFunc("/template", myTemplate)
  mytemp = template.Must(template.ParseFiles("form.html", "linkTemplate.html", "template.html"))
}

func executeMyTemplate (w http.ResponseWriter, template string, data User) {
  err := mytemp.ExecuteTemplate(w, template+".html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  executeMyTemplate(w, "form", User{});
}

func formhandler(w http.ResponseWriter, r *http.Request) {
  myUser := User {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")};
  executeMyTemplate(w, "linkTemplate", myUser);
}

func myTemplate(w http.ResponseWriter, r *http.Request) {
  myUser := User {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")};
  executeMyTemplate(w, "template", myUser);
}