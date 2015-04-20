package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/results", showResults)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	rootForm, err := ioutil.ReadFile("pages/rootForm.html");
	if err != nil {
		fmt.Fprint(w, "404 Not Found")
	}
	fmt.Fprint(w, string(rootForm))
}

var results, _ = ioutil.ReadFile("pages/results.html");
var resultsTemplate = template.Must(template.New("results").Parse(string(results)))

func showResults(w http.ResponseWriter, r *http.Request) {
	strEntered := r.FormValue("str")
	if strEntered == "Jenni" {
		err := resultsTemplate.Execute(w, "Congratulations! You know your own name.")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		err := resultsTemplate.Execute(w, "You filthy little liar!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

