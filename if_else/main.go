package main

import (
	"fmt"
	"math/rand"
)

func main() {
	if true {
		fmt.Println("Condition is true.")
	} else {
		fmt.Println("Condition is false.")
	}

	score := rand.Intn(101)

	if score < 60 {
		fmt.Println("E")
	} else if 60 <= score && score <= 69 {
		fmt.Println("D")
	} else if 70 <= score && score <= 79 {
		fmt.Println("C")
	} else if 80 <= score && score <= 89 {
		fmt.Println("B")
	} else if 90 <= score && score <= 100 {
		fmt.Println("A")
	}
}
