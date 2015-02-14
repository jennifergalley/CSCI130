package main

import "fmt"

func main () {

	//for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//while loop
	j := 0
	for j < 10 {
		fmt.Println (j)
		j++
	}

	//use break
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println (i)
	}

	//use continue
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		} else {
			fmt.Println (i)
		}
	}

	//declare a slice
	mySlice := []string { //like an array
		"Hello",
		"David",
	}

	//loop over a slice
	for i, val := range mySlice {
		fmt.Println (i, " - ", val)
	}

	//declare a map
	myMap := map[string]string { //like an associative array
		"name":"Jenni",
		"age":"20",
		"major":"CSCI",
		"grade":"Junior",
	}

	//loop over a map
	for key, val := range myMap {
		fmt.Println ("Key: ", key, " Value: ", val)
	}


}