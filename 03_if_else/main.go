package main

import (
	"fmt"
	"math"
)

func main() {
	score := 80
	if score >= 80 && score <= 100 {
		fmt.Println("Score: ", score, "Grade: A 😀")
	} else if score >= 70 && score < 80 {
		fmt.Println("Score: ", score, "Grade: B 😄")
	} else if score >= 60 && score < 70 {
		fmt.Println("Score: ", score, "Grade: C 😁")
	} else if score >= 50 && score < 60 {
		fmt.Println("Score: ", score, "Grade: D 😆")
	} else {
		fmt.Println("Score: ", score, "Grade: F 😅")
	}

	// Short if
	limit := 225.0
	if v := math.Pow(10, 2); v < limit {
		fmt.Println("x power n is:", v)
	} else {
		fmt.Printf("x power n is %g over limit %g.\n", v, limit)
	}
}
