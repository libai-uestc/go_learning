package main

import "fmt"

func main() {
	if 5 > 9 {
		fmt.Println("AAA")
	}
	var a int = 10
	if a < 5 {
		fmt.Println("BBB")
	}
	// 在if表达式里面还可以创建一个局部变量
	if b := 8; b > a {
		fmt.Println("b>a")
	} else {
		fmt.Println("b<=a")
	}
	if c, d, e := 1, 4, 7; c < d && (c > e || c > 3) {
		fmt.Println("DDD")
	} else if 3 < e {
		fmt.Println("3 < e")
	} else {
		fmt.Println("hahaha!!!")
	}

}
