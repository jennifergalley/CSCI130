package spaghetti

import (
    "fmt"
    "html/template"
    "net/http"
    "time"

    "appengine"
    "appengine/datastore"
    "appengine/user"
)

type Event struct {
    Title string
    Datetime string
    Details string
}

func init () {
    //handle static paths - do this in app.yaml now
    //http.Handle("/static/img/", http.StripPrefix("/static/img/", http.FileServer(http.Dir("static/img"))))
    
    //handle application paths
    http.HandleFunc("/admin/events/add", addEventHandler)
    http.HandleFunc("/events", eventHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/", rootHandler)
}

//error handler function
func errorHandler(w http.ResponseWriter, r *http.Request, status int, err string) {
    w.WriteHeader(status)
    switch (status) {
    case http.StatusNotFound:
        page := template.Must(template.ParseFiles(
            "static/_base.html",
            "static/404.html",
        ))

        if err := page.Execute(w, nil); err != nil {
            errorHandler(w, r, http.StatusInternalServerError, err.Error())
            return
        }
    case http.StatusInternalServerError:
        page := template.Must(template.ParseFiles(
            "static/_base.html",
            "static/500.html",
        ))

        if err := page.Execute(w, nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}

func eventList(c appengine.Context) *datastore.Key {
    return datastore.NewKey(c, "Events", "default_eventlist", 0, nil)
}

func addEventHandler(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r) //get context
    u := user.Current(c) //get user info
    if u == nil { //not logged in - force login
        url, err := user.LoginURL(c, r.URL.String()) //have user login
        if err != nil {
            errorHandler(w, r, http.StatusInternalServerError, err.Error())
            return
        }
        w.Header().Set("Location", url) //set location for redirect
        w.WriteHeader(http.StatusFound)
        return
    }

    //check the request method
    if r.Method == "POST" {
        //handle post requests
        g := Event{ //create an Event object
            Title: r.FormValue("title"), //populate Event with form data
            Datetime: r.FormValue("date"),
            Details: r.FormValue("details"),
        }

        key := datastore.NewIncompleteKey(c, "Events", eventList(c)) //return index key
        _, err := datastore.Put(c, key, &g) //write to datastore
        
        //error handling
        if err != nil {
            errorHandler(w, r, http.StatusInternalServerError, err.Error())
            return
        }
        http.Redirect(w, r, "/events", http.StatusFound) //redirect back to events page
    } else if r.Method == "GET" { 
        //handle get requests
        page := template.Must(template.ParseFiles( //send out html
            "static/_base.html",
            "static/admin/add-event.html",
        ))

        //error handling
        if err := page.Execute(w, nil); err != nil {
            errorHandler(w, r, http.StatusInternalServerError, err.Error())
            return
        }
    } else {
        fmt.Fprint(w, r.Method)
    }
}

func eventHandler (w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r) //pull in context
    //make query from Events table, allow parent-child relationships, order by datetime descending, limit to 10 results
    q := datastore.NewQuery("Events").Ancestor(eventList(c)).Order("-Datetime").Limit(10) 
    events := make([]Event, 0, 10) //make a slice
    if _, err := q.GetAll(c, &events); err != nil { //pass data from query into slice
        errorHandler(w, r, http.StatusInternalServerError, err.Error())
        return
    }

    //loop through events, changing events.Datetime from YYY-MM-DDTHH:MM
    //to YYYY-MM-DD {0}:MM
    for key, event := range events {
        t, err := time.Parse("2006-01-02T15:04", event.Datetime)
        if err != nil {
            errorHandler(w, r, http.StatusInternalServerError, err.Error())
            return
        }

        events[key].Datetime = t.Format("2006-01-02 03:04 PM")
    }

    //call up static files
    page := template.Must(template.ParseFiles(
        "static/_base.html",
        "static/events.html",
    ))

    //pass in slice to templates
    if err := page.Execute(w, events); err != nil {
        errorHandler(w, r, http.StatusInternalServerError, err.Error())
        return
    }
}

//handler for the about page
func aboutHandler (w http.ResponseWriter, r *http.Request) {
    page := template.Must(template.ParseFiles(
        "static/_base.html",
        "static/about.html",
    ))

    if err := page.Execute(w, nil); err != nil {
        errorHandler(w, r, http.StatusInternalServerError, err.Error())
        return
    }
}

//w is for outputting to screen, r is the input request variables
func rootHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound, "")
        return
    }

    //build a template, parse these static files
    page := template.Must(template.ParseFiles(
        "static/_base.html",
        "static/index.html",
    ))

    //error checking - execute the page and if error, stop
    if err := page.Execute(w, nil); err != nil {
        panic (err)
    }
}
