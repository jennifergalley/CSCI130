package main

import "fmt"

//Struct
type Friend struct {
	name string
	phone int
}

//Type
type Salutation string

func main () {
	var message1 string = "Hello world!"
	fmt.Print(message1)
	fmt.Println(message1)

	var message2 = "Hello world!"
	fmt.Println(message2)

	var alpha, beta, chai = 1, 2, 3
	fmt.Println(alpha, beta, chai)

	message3 := "Hello world!"
	a, b, c := 1, false, 3
	fmt.Println(message3, a, b, c)

	//Pointer
	var message = "Hello world!"
	var greeting *string = &message
	fmt.Println(message, *greeting)

	var alex = Friend {}
	alex.name = "Alex"
	alex.phone = 123456789
	fmt.Println (alex.name, alex.phone)

	var joe = Friend {"Joe", 123456789}
	fmt.Println (joe.name, joe.phone)

	var chelsea = Friend {name:"Chelsea", phone:123456789}
	fmt.Println (chelsea.name, chelsea.phone)

	var hi Salutation = "hi"
	fmt.Println (hi)
}

//go run src.go