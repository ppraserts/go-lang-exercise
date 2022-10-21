package main

import "fmt"

var msg3 string = "Hello Gopher message 3!!!"

func main() {
	// Declare variable with string
	var msg1 string = "Hello Gopher message 1!!!"
	fmt.Println(msg1)

	// Declare variable without string
	msg2 := "Hello Gopher message 2!!!"
	fmt.Println(msg2)

	// Declare global variable
	// fmt.Println(msg3)

	// Declare variable scope
	{
		msg4 := "Hello Gopher message 4!!!"
		fmt.Println(msg4)
	}

	// Static callytype
	msg5 := "message 5"
	msg5 = "Hello Gopher message 5!!!"
	fmt.Println(msg5)

	// Concatenate string
	msg6 := "Hello Gopher " + "message 6!!!"
	fmt.Println(msg6)

	// Go not allow declared but not used
	// msg7 := "Hello Gopher message 7!!!"

	// Single Comment
	/*
		Multiple Comment
	*/

	// Print format value ... reference: https://pkg.go.dev/fmt#Println
	msg8 := "Hello Gopher "
	fmt.Println(msg8, "message 8!!!")

	// Multiple Declaration
	msg9, age, price, check := "Hello Gopher message 9!!!", 38, 22.52, true
	fmt.Println(msg9, "Age: ", age, " Price: ", price, "Check: ", check)

	// Raw String
	msg10 := `
		Hello Gopher message 10!!!
	`
	fmt.Println(msg10)
}
