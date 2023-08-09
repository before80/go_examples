package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(name string) {
	fmt.Println("Hello,", name)
}

type Student struct {
	RemainCourses []string
	mu            sync.Mutex
}

func (s *Student) DoHomework() {
	if len(s.RemainCourses) == 0 {
		fmt.Println("do nothing...")
		return
	}
	h := s.RemainCourses[0]
	s.mu.Lock()
	if len(s.RemainCourses) == 1 {
		s.RemainCourses = nil
	} else {
		s.RemainCourses = s.RemainCourses[1:]
	}
	s.mu.Unlock()

	fmt.Println("do ", h, "")
}

func main() {

	go func() {
		fmt.Println("In 1 goroutine.")
	}()

	go func() {
		fmt.Println("In 4 goroutine.")
	}()

	go func() {
		fmt.Println("In 3 goroutine.")
	}()

	go func() {
		fmt.Println("In 2 goroutine.")
	}()

	time.Sleep(time.Second)

	s1 := Student{RemainCourses: []string{"Math", "Chinese", "English"}}
	go s1.DoHomework()
	go s1.DoHomework()
	go s1.DoHomework()
	go s1.DoHomework()
	go s1.DoHomework()

	time.Sleep(time.Second)
}
