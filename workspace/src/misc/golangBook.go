package main

import (
	"fmt"
	// "math"
	"time"
)

func fahrenheitToCelsius () {
	fmt.Print("Enter the temperature in Fahrenheit: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := (input - 32) * (5.0/9.0)

	fmt.Println (output)
}

func feetToMeters () {
	fmt.Print("Enter the distance in feet: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 0.3048

	fmt.Println (output)
}


func divisibleBy3 () {
	for i := 1; i <= 100; i++ {
       if i % 3 == 0 {
            fmt.Println(i)
        }
    }
}


func fizzBuzz () {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println ("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println ("Fizz")
		} else if i % 5 == 0 {
			fmt.Println ("Buzz")
		} else {
			fmt.Println (i)
		}
	}
}


func smallestInList () int {
	x := []int{ 48,96,86,68, 57,82,63,70, 37,34,83,27, 19,97, 9,17, }
	min := x[0]
	for i := 1; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

func half (x int) (int, bool) {
	return (x/2), (x%2 == 0)
}

func max (x ...int) int {
	max := x[0]
	for _, v := range x {
		if v > max {
			max = v
		}
	}
	return max
}

func makeOddGenerator() func() uint {
    i := uint(1)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}

func fib (n int) int {
	switch n {
		case 0:
			return 0
		case 1:
			return 1
		default: 
			return fib(n-1) + fib(n-2) 
	}
}

func swap (x *int, y *int) {
	temp := *y
	*y = *x
	*x = temp
}

/*type Circle struct {
	x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func (c *Circle) perimeter () float64 {
	return math.Pi * c.r * 2
}

func distance (x1, y1, x2, y2 float64) float64 {
    //copy and paste doesn't work exactly from the website. because it changes "-"
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) perimeter () float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * 2 + w * 2
}

type Shape interface {
	perimeter() float64
}*/

func sleep (n int) {
	for i := 0; i < n; i++ {
		<- time.After(time.Second)
	}
}

func Min (x []float64) float64 {
	min := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] < min {
			min = x[i]
		}
	}
	return min
}

func Max (x []float64) float64 {
	max := x[0]
	for i := 0; i < len(x); i++ {
		if x[i] > max {
			max = x[i]
		}
	}
	return max
}

func main () {
	// fahrenheitToCelsius ()
	// feetToMeters ()
	// divisibleBy3 ()
	// fizzBuzz ()
	// fmt.Print (smallestInList ())
	// fmt.Print (half(6))
	// fmt.Print (half(5))
	// fmt.Print (max(1, 2, 3, 4))
	/*nextOdd := makeOddGenerator ()
	fmt.Println (nextOdd())
	fmt.Println (nextOdd())
	fmt.Println (nextOdd())*/
	// fmt.Print (fib(10))
	/*x := 1
	y := 2
	swap (&x, &y)
	fmt.Print (x, y)*/
	/*c := Rectangle {4, 5}
	fmt.Print(c.perimeter())*/
	/*sleep (3)
	fmt.Print("done")*/
	/*s := []float64 {1.0, 2.0, 3.0}
	fmt.Println (Min(s))
	fmt.Println (Max(s))*/
}
