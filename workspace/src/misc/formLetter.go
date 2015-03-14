package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Define a template.
	const letter = `
Dear {{.Honorific}} {{.LastName}},
{{if .Attended}}
It was a pleasure to see you at the fundraiser.{{else}}
It is a shame you couldn't make it to the fundraiser.{{end}}
{{with .Donated}}Thank you for the generous donation.{{end}}
Don't forget about our upcoming events!
{{range .Events}}
	{{.}}
{{end}}
Sincerely,
Ms. Garner
`

	// Prepare some data to insert into the template.
	type Client struct {
		LastName, Honorific string
		Attended bool
		Donated int
		Events []string
	}
	var events = []string {"Founder's Day Parade", "Awards Night Gala", "Independence Day Ball"}
	var clients = []Client{
		{"Gilbert", "Mr.", true, 0, events},
		{"Johnson", "Mrs.", false, 0, events},
		{"Lockwood", "Mr.", true, 1000, events},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range clients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

}
