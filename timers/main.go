// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	fmt.Println("Waiting...")
	<-timer1.C
	fmt.Println("timer1 expired")

	_ = time.AfterFunc(2*time.Second, func() {
		fmt.Println("timer2 expired")
	})

	fmt.Println("Waiting...")
	time.Sleep(2010 * time.Millisecond) // 2.01s > 前面的 2s

	fmt.Println("End")
}
