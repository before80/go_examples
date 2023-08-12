// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting at %s\n", id, time.Now())
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("Worker %d done at %s\n", id, time.Now())
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
