// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now) // 2023-08-10 16:46:30.7978556 +0800 CST m=+0.005433901

	fmt.Println(now.Unix())      // 1691657190
	fmt.Println(now.UnixMilli()) // 1691657190797
	fmt.Println(now.UnixNano())  // 1691657190797855600

	fmt.Println(time.Unix(now.Unix(), 0))     // 2023-08-10 16:46:30 +0800 CST
	fmt.Println(time.Unix(0, now.UnixNano())) // 2023-08-10 16:46:30.7978556 +0800 CST
}
