package main

import (
	"html/template"
	"net/http"
    "time"
)

var mytemp *template.Template

type User struct{
    Name string
    Email string
    Message string
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/linkTemplate", linkTemplate)
	http.HandleFunc("/messageTemplate", messageTemplate)
	http.HandleFunc("/cookieTemplate", cookieTemplate)
    mytemp = template.Must(template.ParseFiles("form.html","linkTemplate.html","messageTemplate.html","cookieTemplate.html"))
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

func linkTemplate(w http.ResponseWriter, r *http.Request) {
    expire := time.Now().AddDate(0, 0, 1)

    myUser := User{
        Name: r.FormValue("name"),
        Email: r.FormValue("email"),
        Message: r.FormValue("message"),
    }

    cookie := http.Cookie{
        Name:       "test",
        Value:      myUser.Message,
        Expires:    expire,
        RawExpires: expire.Format(time.UnixDate),
        MaxAge:   86400,
        Secure:   false,
        HttpOnly: false,
    }
    http.SetCookie(w, &cookie)

    executeMyTemplate(w, "linkTemplate", myUser)
}

func messageTemplate(w http.ResponseWriter, r *http.Request) {
    myUser := User {
        r.FormValue("name"), 
        r.FormValue("email"), 
        r.FormValue("message"),
    };
    executeMyTemplate(w, "messageTemplate", myUser)
}

func cookieTemplate(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("test")
    myUser := User {"", "", cookie.Value}
    executeMyTemplate(w, "cookieTemplate", myUser);
}