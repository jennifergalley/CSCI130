package main

import "fmt"

type Person struct {
	greeting string
	name     string
}

type myPrintType func(string)

func Greet(person Person, foobar func(string)) {
	g, n := CreateMessage(person.name, person.greeting, "howdy")
	foobar(g)
	foobar(n)
}

func CreateMessage(name string, greeting ...string) (myGreeting string, myName string) {
	myGreeting = greeting[0] + " " + name
	myName = "\nHey, " + name + "\n"
	return
}

func myPrint(s string) {
	fmt.Print(s)
}

func myPrintFunction(custom string) myPrintType {
	return func(s string) {
		fmt.Print(s + custom)
	}
}

func main() {
	v := Person{}
	v.greeting = "Why are you here,"
	v.name = "Jenni?"
	Greet(v, myPrintFunction("^^^"))

	if (true) {
		fmt.Println ("Hello, world")
	}

	if (3 > 5) {
		fmt.Println ("This is ran true")
	} else {
		fmt.Println ("This ran false")
	}

	myvar := 5
	switch (myvar) {
	case 0:
		fmt.Println (0)
	case 1:
		fmt.Println (1)
	case 2:
		fmt.Println (2)
	case 3:
		fmt.Println (3)
	case 4:
		fmt.Println (4)
	case 5:
		fmt.Println (5)
	}


}