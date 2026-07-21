package main

import "fmt"

func slice_init() {
	var s []int // 切片的[]里为空，数组必须要填值或者[...]
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{}
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3) // len = cap = 3 make关键字，第一个参数是切片（类型），第二个参数是长度，若无第三个参数（容量），则默认长度=容量
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = make([]int, 3, 5) //len = 3 cap = 5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	s = []int{1, 2, 3, 4, 5} // len = 5 cap = 5
	fmt.Printf("len %d cap %d\n", len(s), cap(s))
	fmt.Println("--------------------------------------")

	s2d := [][]int{
		{1},
		{2, 3},
	}
	fmt.Printf("s2d len %d cap %d\n", len(s2d), cap(s2d))
	fmt.Printf("s2d[0] len %d cap %d\n", len(s2d[0]), cap(s2d))
	fmt.Printf("s2d[1] len %d cap %d\n", len(s2d[0]), cap(s2d[0]))
	fmt.Println("--------------------------------------")
}
