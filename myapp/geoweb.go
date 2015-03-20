package main

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/directions", directions)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Google Directions API</title>
      </head>
      <body>
      <h1>Get Directions</h1>
        <form style="margin-left: 120px;" action="/directions" method="post" accept-charset="utf-8">
        <label for="str1">Origin Address</label><br/><br/>
          <input type="text" name="str1" placeholder="Type address..." id="str" /><br/><br/>
        <label for="str2">Destination Address</label><br/><br/>
          <input type="text" name="str2" placeholder="Type address..." id="str" /><br/><br/>
          <input type="submit" value="Submit" />
        </form>
      </body>
    </html>
`

var upperTemplate = template.Must(template.New("directions").Parse(upperTemplateHTML))

func directions(w http.ResponseWriter, r *http.Request) {
	// Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
	addr1 := r.FormValue("str1")
	addr2 := r.FormValue("str2")

	safeAddr1 := url.QueryEscape(addr1)
	safeAddr2 := url.QueryEscape(addr2)
	fullUrl := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/directions/json?origin=%s&destination=%s&key=AIzaSyAgIIyyQNVy3B583nr5_AotpEycftScxFg", safeAddr1, safeAddr2)

	// Build the request
	req, err := http.Get(fullUrl)
	if err != nil {
		panic(err)
	}

	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
	     panic(err)
	}

	tempErr := upperTemplate.Execute(w, string(jsonDataFromHttp))
	if tempErr != nil {
		http.Error(w, tempErr.Error(), http.StatusInternalServerError)
	}
}

const upperTemplateHTML = `
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Display Image</title>
      <link rel="stylesheet" href="/stylesheets/goview.css">
    </head>
    <body>
      {{html .}}
    </body>
  </html>
`
