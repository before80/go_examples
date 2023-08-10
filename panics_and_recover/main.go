package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%v,%T\n", r, r)
		}
	}()

	panic("a problem")
}

// Output:
// a problem,string
