package guestbook

import (
    "html/template"
    "net/http"
    "time"
    "strconv"
    "strings"
    // "fmt"
    "appengine"
    "appengine/datastore"
    "appengine/user"
)

type Entry struct {
    G Greeting
    K int64
}

type Greeting struct {
    Author  string
    Content string
    Date    time.Time
}


/*
**************************
Global Variables
**************************
*/
var viewTemplate = template.Must(template.ParseFiles("view.html"))
var editTemplate = template.Must(template.ParseFiles("edit.html"))
var mineTemplate = template.Must(template.ParseFiles("myPosts.html"))


/*
**************************
Init
**************************
*/
func init () {
    http.HandleFunc ("/", redir)
}

func redir(w http.ResponseWriter, r *http.Request) {
    if (strings.HasPrefix(r.URL.Path, "/sign")) {
        sign(w,r)
    }else if(strings.HasPrefix(r.URL.Path, "/edit")){
        edit(w,r)
    }else if(strings.HasPrefix(r.URL.Path, "/submitEdit")){
        submitEdit(w,r)
    }else if(strings.HasPrefix(r.URL.Path, "/delete")){
        delete(w,r)
    }else if (strings.HasPrefix(r.URL.Path, "/mine")) {
        mine(w,r)
    }else if(r.URL.Path == "/"){
        view(w, r)
    }else {
        static(w,r)
    }
}


/*
**************************
Useful Functions
**************************
*/

func guestbookKey(c appengine.Context) *datastore.Key {
    return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}


/* 
****************************
Handlers 
****************************
*/

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "public/" + r.URL.Path)
}

// Upon submit of form, send entry to datastore
func sign(w http.ResponseWriter, r *http.Request) {
    // Write new entry to datastore
    c := appengine.NewContext(r)
    g := Greeting{
            Content: r.FormValue("content"),
            Date:    time.Now(),
    }

    if u := user.Current(c); u != nil {
            g.Author = u.String()
    }

    key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
    _, err := datastore.Put(c, key, &g)

    if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    }

    http.Redirect(w, r, "/mine", http.StatusFound)
}

// Fetch and display Datastore objects
func view(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    // Query for and display previous entries
    q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
    greetings := make([]Greeting, 0, 10)
    keys, err := q.GetAll(c, &greetings)
    // fmt.Fprint(w, greetings)
    // fmt.Fprint(w, keys)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    entries := make([]Entry, len(keys))
    for i := 0; i < len(keys); i++ {
        entries[i] = Entry{G: greetings[i], K: keys[i].IntID()}
    }
    // fmt.Fprint(w, entries)

    if err := viewTemplate.Execute(w, entries); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Fetch and display Datastore objects
func mine(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    var loggedIn = ""
    if u := user.Current(c); u != nil {
        loggedIn = u.String()
    }
    // fmt.Fprint(w, loggedIn)

    // Query for and display previous entries
    q := datastore.NewQuery("Greeting").Filter("Author =", loggedIn)
    greetings := make([]Greeting, 0, 10)
    keys, err := q.GetAll(c, &greetings)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    entries := make([]Entry, len(keys))
    for i := 0; i < len(keys); i++ {
        if (greetings[i].Author == loggedIn){
            entries[i] = Entry{G: greetings[i], K: keys[i].IntID()}
        }
    }
    // fmt.Fprint(w, entries)

    if err := mineTemplate.Execute(w, entries); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// We want to edit an entry - take us to the edit form
func edit(w http.ResponseWriter, r *http.Request) {
    k, _ := strconv.Atoi(r.FormValue("key"))

    g := Greeting{
        Content: r.FormValue("content"),
    }

    entry := Entry{G: g, K: int64(k),}

    if err := editTemplate.Execute(w, entry); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Upon submit of edit form, change entry in datastore
func submitEdit(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

    k, _ := strconv.Atoi(r.FormValue("key"))

    g := Greeting{
            Content: r.FormValue("content"),
            Date:    time.Now(),
    }

    if u := user.Current(c); u != nil {
            g.Author = u.String()
    }

    key := datastore.NewKey(c, "Greeting", "", int64(k), guestbookKey(c))
    _, err := datastore.Put(c, key, &g)

    if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    }

    http.Redirect(w, r, "/mine", http.StatusFound)
}

func delete (w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
    k, _ := strconv.Atoi(r.FormValue("key"))
    key := datastore.NewKey(c, "Greeting", "", int64(k), guestbookKey(c))
    if err := datastore.Delete(c, key); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    http.Redirect(w, r, "/mine", http.StatusFound)
}
