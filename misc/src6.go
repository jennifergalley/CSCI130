package main

import "fmt"

type MySentence string

type Contact struct {
	name string
	greeting string
}

func (c *Contact) rename (newName string) { //to change struct property
	c.name = newName
}

func (s MySentence) eatChocolate () {
	fmt.Println ("METHOD: EAT MORE CHOCOLATE NOW!")
}

func (s MySentence) drinkMilk () {
	fmt.Println ("METHOD: DRINK MILK TO GO WITH THAT CHOCOLATE")
}

func main () {
	var message MySentence = "Hello world!"
	fmt.Println ("DATA: " + message)
	message.eatChocolate()
	message.drinkMilk ()

	marcus := Contact{"Marcus", "Hello!"}
	fmt.Println(marcus)
	marcus.rename("Jenny")
	fmt.Println(marcus)
}