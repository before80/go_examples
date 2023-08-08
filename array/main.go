package main

import "fmt"

func main() {
	// 一维数组
	a1 := [3]int{1, 2, 3}

	// 二维数组
	a2 := [3][3]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	// 三维数组
	a3 := [3][3][3]int{
		{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
		{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
		{
			{1, 2, 3},
			{1, 2, 3},
			{1, 2, 3},
		},
	}

	fmt.Println("--------- %#v ----------")
	fmt.Printf("%#v\n", a1)
	fmt.Printf("%#v\n", a2)
	fmt.Printf("%#v\n", a3)

	fmt.Println("--------- %+v ----------")
	fmt.Printf("%+v\n", a1)
	fmt.Printf("%+v\n", a2)
	fmt.Printf("%+v\n", a3)

	fmt.Println("--------- %v ----------")
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
	fmt.Printf("%v\n", a3)
}
