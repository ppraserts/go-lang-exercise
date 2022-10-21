package main

import "fmt"

func main() {
	// Short switch-case
	switch today := "Satureday"; today {
	case "Satureday", "Sunday":
		fmt.Println(today, "is Weekend 😍.")
	default:
		fmt.Println(today, "is Working Day 😅.")
	}

	// switch-case with no expression
	num := 1
	switch {
	case num < 0:
		fmt.Println("nagative number.😆")
	case num == 0:
		fmt.Println("zero number.😀")
	case num > 0:
		fmt.Println("positive number.😍")
	default:
		fmt.Println("😅")
	}
}
