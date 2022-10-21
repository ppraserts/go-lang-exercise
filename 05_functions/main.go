package main

import "fmt"

//  Multiple return values
func add(x, y float64) (float64, string) {
	fmt.Println(x, "+", y, "=", x+y)
	return x + y, "Result from add functions:"
}

// Naked returns
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// First Class function ---> Assigning a function to a variable
var addFunc = func(x, y float64) float64 {
	return x + y
}

// First Class function ---> Passing a function as an argument
func compute(fn func(float64, float64) float64) float64 {
	v := fn(42, 13)
	return v
}

// Higher-order function ---> Returning a function
func adder() (func() int, func() int) {
	sum := 0
	return func() int {
			sum = sum + 1
			return sum
		}, func() int {
			return sum
		}
}

func main() {
	result, message := add(42, 13)
	fmt.Println(message, result)

	n1, n2 := split(10)
	fmt.Println("n1:", n1, "n2:", n2)

	fmt.Println("First Class function ---> Assigning a function to a variable: ", addFunc(42, 13))
	fmt.Printf("%T\n", addFunc) // print to show signature function

	fmt.Println("First Class function ---> Passing a function as an argument: ", compute(addFunc))

	i, c := adder()
	fmt.Println("higher-order function ---> call increae: ", i())
	fmt.Println("higher-order function ---> call increae: ", i())
	fmt.Println("higher-order function ---> get current: ", c())
}
