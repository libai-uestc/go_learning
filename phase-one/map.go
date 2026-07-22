package main

import "fmt"

func key_type() {
	// 函数类型、map类型、切片以及包含这三种类型的结构体不能作为key
	// map的value可以是任意类型
	type f func(int) bool
	type m map[int]byte
	type s []int

	type i int

	var m1 map[i]f
	fmt.Println(m1)

	type user struct {
		scores float32 //如果scores是slice，则user不能作为map的key
	}
	u := user{}
	m5 := make(map[user]interface{})
	m5[u] = 5
	fmt.Println(m5)

	var ft float32
	m6 := make(map[float32]any)
	m6[ft] = 9
}

func update_map() {
	var m map[string]int
	m = make(map[string]int)
	m = make(map[string]int, 5)
	m = map[string]int{"语文": 0, "数学": 39, "物理": 57, "历史": 49}
	m["英语"] = 100
	fmt.Println(m["数学"])
	delete(m, "数学")
	fmt.Println(m["数学"])
	if value, exists := m["语文"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("map里不存在[语文]这个key")
	}
	// 可以获取map的len，无法获取cap
	fmt.Printf("map里有%d对KV\n", len(m))
	// 遍历map
	for key, value := range m {
		fmt.Printf("%s=%d\n", key, value)
	}
	fmt.Println("以上是第一次遍历-----------")
	for key, value := range m {
		fmt.Printf("%s=%d\n", key, value)
	}
	fmt.Println("以上是第二次遍历-----------")

	// 一边遍历一边修改
	for key, value := range m {
		m[key] = value + 1
	}
	for key, value := range m {
		fmt.Printf("%s=%d\n", key, value)
	}
	fmt.Println("---------------")

	// for range获取的值是拷贝
	for _, value := range m {
		value = value + 1
	}
	for key, value := range m {
		fmt.Printf("%s=%d\n", key, value)
	}
}

func main() {
	key_type()
	update_map()
}
