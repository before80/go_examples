package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Printf("%#v,%T\n", argsWithProg, argsWithProg)
	fmt.Printf("%v,%T\n", argsWithoutProg, argsWithoutProg)
	fmt.Printf("%v,%T\n", arg, arg)
}
