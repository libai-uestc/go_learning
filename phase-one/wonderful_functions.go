package main

import (
	"fmt"
	"slices"
)

func main21() {
	arr := []int{3, 7, 4, 6, 8, 1}
	slices.Sort(arr)
	fmt.Println(arr)
	slices.SortFunc(arr, func(a, b int) int {
		return a - b //前者减后者，就是从小到大；后者减前者，就是从大到小
	})
	fmt.Println(arr)

	type User struct {
		Age    int
		Height float32
	}
	brr := []*User{{18, 1.8}, {25, 1.7}}
	slices.SortFunc(brr, func(a, b *User) int {
		if b.Height > a.Height {
			return 1
		} else if b.Height < a.Height {
			return -1
		} else {
			return 0
		}
	})

	fmt.Println("最大者", slices.Max(arr))
	fmt.Println("最小者", slices.Min(arr))
	fmt.Println("包含", slices.Contains(arr, 5))

	crr := make([]int, len(arr))
	copy(crr, arr)
	fmt.Println(crr)

	fmt.Println("相等", slices.Equal(arr, crr)) //只比较值是否相等，true
	arr[0]++
	fmt.Println("相等", slices.Equal(arr, crr)) //false

	drr := brr                                //共享底层数组空间
	fmt.Println("相等", slices.Equal(brr, drr)) //true
	arr[0]++
	fmt.Println("相等", slices.Equal(brr, drr)) //true
}
