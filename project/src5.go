package main

import (
	"fmt"
	"myMath"
)

type Intern struct {
	name string
	age int
	company string
}

type MyString string

func (i Intern) apply () {
	fmt.Println ("Applicant ", i.name, ", age ", i.age, " is applying to ", i.company)
}

func (s MyString) print () {
	fmt.Println (s)
}

func (i Intern) accepted () bool {
	if i.name == "Jenni" && i.age == 20 && i.company == "Microsoft" {
		return true
	} else {
		return false
	}
}

func main () {

	//append to a slice
	mySlice := []string {"Hello", "world!"}
	mySlice = append(mySlice, "Bonjour!")
	fmt.Println(mySlice)

	//delete from a slice
	mySlice = append(mySlice[:2])
	fmt.Println(mySlice)

	//compare a slice and an array
	fmt.Print ("This is an array: ")
	myArray := [5] int {0, 1, 2, 3, 4}
	fmt.Print (myArray)
	fmt.Println ()

	fmt.Print ("This is a slice: ")
	mySlice3 := [] int {0, 1, 2, 3, 4}
	mySlice3 = append (mySlice3, 5)
	fmt.Print (mySlice3)
	fmt.Println ()

	//packages
	div := myMath.Divide (15, 3)
	fmt.Println(div)

	//methods attached to a struct
	var jenni = Intern {}
	jenni.name = "Jenni"
	jenni.age = 20
	jenni.company = "Microsoft"
	jenni.apply()

	//methods attached to a string
	var pumpkin MyString = "pumpkin"
	pumpkin.print ()

	//methods with return values
	fmt.Println (jenni.accepted())
}