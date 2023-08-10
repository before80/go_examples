package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type A struct {
	done atomic.Value
	mu   sync.Mutex
}

func (a *A) Done() <-chan int {
	d := a.done.Load()
	if d != nil {
		return d.(chan int)
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	d = a.done.Load()
	if d == nil {
		d = make(chan int)
		a.done.Store(d)
	}
	return d.(chan int)
}

var closedchan = make(chan int)

func init() {
	close(closedchan)
}
func main() {
	a := A{}
	a.done.Store(make(chan int))

	fmt.Printf("done=%v,%T\n", a.done.Load(), a.done.Load())

	fmt.Printf("Done()=%v,%T\n", a.Done(), a.Done())
	fmt.Printf("Done()=%v,%T\n", a.Done(), a.Done())

	//d, _ := a.done.Load().(chan int)
	//close(d)
	a.done.Store(closedchan)

	fmt.Printf("done=%v,%T\n", a.done.Load(), a.done.Load())

	fmt.Printf("Done()=%v,%T\n", a.Done(), a.Done())
	fmt.Printf("Done()=%v,%T\n", <-a.Done(), <-a.Done())
}
