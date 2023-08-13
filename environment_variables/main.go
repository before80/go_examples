// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO")) // FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR:

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
