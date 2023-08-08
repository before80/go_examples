package main

import "fmt"

// 无类型常量
const B10 = false
const I10 = 1
const F10 = 1.2

// 有类型常量
const B20 bool = true
const I20 int8 = 1
const F20 float32 = 1.2

// 使用iota
const (
	AI = iota + 1
	BI
	CI
	_
	DI
)

func main() {
	fmt.Printf("%v,%T\n", B10, B10)
	fmt.Printf("%v,%T\n", I10, I10)
	fmt.Printf("%v,%T\n", I10, I10)

	fmt.Printf("%v,%T\n", B20, B20)
	fmt.Printf("%v,%T\n", I20, I20)
	fmt.Printf("%v,%T\n", F20, F20)

	fmt.Printf("%v,%T\n", AI, AI)
	fmt.Printf("%v,%T\n", BI, BI)
	fmt.Printf("%v,%T\n", CI, CI)
	fmt.Printf("%v,%T\n", DI, DI)

	// 无类型常量
	const b10 = false
	const i10 = 1
	const f10 = 1.2

	// 有类型常量
	const b20 bool = true
	const i20 int8 = 1
	const f20 float32 = 1.2

	// 使用iota
	const (
		ai = iota + 1
		bi
		ci
		_
		di
	)

	fmt.Printf("%v,%T\n", b10, b10)
	fmt.Printf("%v,%T\n", i10, i10)
	fmt.Printf("%v,%T\n", f10, f10)

	fmt.Printf("%v,%T\n", b20, b20)
	fmt.Printf("%v,%T\n", i20, i20)
	fmt.Printf("%v,%T\n", f20, f20)

	fmt.Printf("%v,%T\n", ai, ai)
	fmt.Printf("%v,%T\n", bi, bi)
	fmt.Printf("%v,%T\n", ci, ci)
	fmt.Printf("%v,%T\n", di, di)
}

// Output:
//false,bool
//1,int
//1,int
//true,bool
//1,int8
//1.2,float32
//1,int
//2,int
//3,int
//5,int
//false,bool
//1,int
//1.2,float64
//true,bool
//1,int8
//1.2,float32
//1,int
//2,int
//3,int
//5,int
