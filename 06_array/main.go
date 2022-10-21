package main

import "fmt"

func main() {
	// skills := [3]string{"C#", "Go", "Python"}
	// or
	skills := [...]string{"C#", "Go", "Python"}
	// or
	// var skills [3]string = [3]string{"C#", "Go", "Python"}
	fmt.Printf("%#v\n", skills)

	l := len(skills)
	fmt.Printf("Skills count: %#v\n", l)

	skills[1] = "Golang"
	fmt.Printf("%#v\n", skills)

	fmt.Printf("%s\n", skills[1])
}
