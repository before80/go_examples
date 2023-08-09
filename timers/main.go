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
	time.Sleep(3 * time.Second)

	fmt.Println("End")
}
