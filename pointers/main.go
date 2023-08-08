package main

import "fmt"

// 不是使用指针，仅仅改变 iVal 变量的值
func setWithVal(iVal, newVal string) {
	iVal = newVal
}

// 使用指针，可以改变 iPtr 所指向的变量的值
func setWithPtr(iPtr *string, newVal string) {
	*iPtr = newVal
}

func main() {
	name := "longX"
	fmt.Println("Initial value is ", name)

	setWithVal(name, "longX-1")
	fmt.Println("After with setWithVal function, name's value is ", name)

	setWithPtr(&name, "longX-2")
	fmt.Println("After with setWithPtr function, name's value is ", name)

}
