package main

import "fmt"

func main() {
	// %s: string, %d: decimal, %f: float, $t: boolean
	msg, age, price, check := "Hello Gopher!!!", 55, 22.52, true
	fmt.Printf("msg: %s\n", msg)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("price: %.2f\n", price)
	fmt.Printf("check: %t\n", check)
	fmt.Printf("-------------------------------\n")
	/*
		Output:
			msg: Hello Gopher!!!
			age: 55
			price: 22.52
			check: true
	*/

	var r rune = '😆'
	fmt.Println("r: ", r)
	fmt.Printf("r: %c\n", r)
	fmt.Printf("-------------------------------\n")
	/*
		Output:
			r:  128518
			r: 😆
	*/

	// %#v can display all type
	fmt.Printf("msg: %#v\n", msg)
	fmt.Printf("age: %#v\n", age)
	fmt.Printf("price: %#v\n", price)
	fmt.Printf("check: %#v\n", check)
	fmt.Printf("r: %#v\n", r)
	fmt.Printf("-------------------------------\n")
	/*
		Output:
			msg: Hello Gopher!!!
			age: 55
			price: 22.52
			check: true
			r: 128518
	*/

	// %T display type of variable
	fmt.Printf("msg: %T\n", msg)
	fmt.Printf("age: %T\n", age)
	fmt.Printf("price: %T\n", price)
	fmt.Printf("check: %T\n", check)
	fmt.Printf("r: %T\n", r)
	fmt.Printf("-------------------------------\n")
	/*
		Output:
			msg: string
			age: int
			price: float64
			check: bool
			r: int32
	*/
}
