package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()
	done := make(chan bool)
	changeTickerDuration := make(chan bool)

	go func() {
		t := time.After(5 * time.Second)
		for {
			select {
			case <-t:
				done <- true
				return
			case v, ok := <-changeTickerDuration:
				fmt.Printf("%t,%t\n", v, ok)
				if v && ok {
					t = time.After(10 * time.Second)
					ticker.Reset(1 * time.Second)
					fmt.Println("had changed ticker duration")
				}
			}
		}
	}()

	go func() {
		time.Sleep(2 * time.Second)
		r := rand.Intn(100)
		if r > 50 {
			changeTickerDuration <- true
			fmt.Println("heavy task, need to change ticker")
		} else {
			changeTickerDuration <- false
		}
	}()

	for {
		select {
		case <-done:
			fmt.Println("Done!")
			return
		case curTime := <-ticker.C:
			fmt.Println("Current time: ", curTime)
		}
	}
}
