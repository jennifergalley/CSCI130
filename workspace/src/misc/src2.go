package main

import "fmt"

//Struct
type Pizza struct {
	kind string
	price int
}

//Functions
func printPizza (pizza Pizza) {
	fmt.Println (pizza.kind, pizza.price)
}

//Multiple returns
func getPizza (pizza Pizza) (string, int) {
	return pizza.kind, pizza.price
}

//Constants
const (
	PI = 3.14
)

func main () {
	var pizza = Pizza {
		kind: "cheese",
		price: 5,
	}
	printPizza(pizza)

	//Multiple returns without error
	var kind, _ = getPizza(pizza)
	fmt.Println(kind)

	fmt.Println (PI)

}
