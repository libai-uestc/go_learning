package main

import "fmt"

func array1d() {
	var arr1 [5]int = [5]int{}
	var arr2 = [5]int{}
	var arr3 = [5]int{3, 2}
	var arr4 = [5]int{2: 15, 4: 30}
	var arr5 = [...]int{3, 2, 6, 5, 4}
	var arr6 = [...]struct {
		name string
		age  int
	}{{"Tom", 18}, {"Jim", 20}}
	fmt.Printf("arr1=%#v\n", arr1)
	fmt.Printf("arr2=%#v\n", arr2)
	fmt.Printf("arr3=%#v\n", arr3)
	fmt.Printf("arr4=%#v\n", arr4)
	fmt.Printf("arr5=%#v\n", arr5)
	fmt.Printf("arr6=%#v\n", arr6)

	fmt.Printf("arr5[0]=%d", arr5[0])
	fmt.Printf("arr5[len(arr5)-1]=%d\n", arr5[len(arr5)-1]) //获取最后一个元素的唯一方法
	fmt.Printf("数组的地址：%p\n", &arr5)
	fmt.Printf("第一个元素的地址：%p\n", &arr5[0])
	fmt.Printf("第二个元素的地址：%p\n", &arr5[1])

	for i, ele := range arr5 {
		fmt.Printf("index=%d element=%d\n", i, ele)
	}
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("index=%d,element=%d\n", i, arr5[i])
	}
	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("cap(arr1)=%d\n", cap(arr1))
}

func array2d() {
	var arr1 = [5][3]int{{1}, {2, 3}}
	var arr2 = [...][3]int{{1}, {2, 3}}

	fmt.Printf("arr[1][1]=%d\n", arr1[1][1])
	fmt.Printf("arr[4][2]=%d\n", arr1[4][2])

	for row, array := range arr2 {
		for col, ele := range array {
			fmt.Printf("arr2[%d][%d]=%d\n", row, col, ele)
		}
	}

	fmt.Printf("len(arr1)=%d\n", len(arr1))
	fmt.Printf("cap(arr1)=%d\n", cap(arr1))
}

func update_array1(arr [5]int) {
	fmt.Printf("array in function, address is %p\n", &arr[0])
	arr[0] = 888
}

// 传数组的指针
func update_array2(arr *[5]int) {
	fmt.Printf("array in function, address is %p\n", &((*arr)[0]))
	arr[0] = 888
}

// 传指针构成的数组
func update_array3(arr [5]*int) {
	*arr[0] = 888
}

func for_range_array() {
	arr := [...]int{1, 2, 3}
	for i, ele := range arr {
		arr[i] += 8
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
		ele += 1
		fmt.Printf("%d %d %d\n", i, arr[i], ele)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d\n", i, arr[i])
	}
}
