package main

import (
	"fmt"
	"time"
)

func main20() {
	c, d := 3, 5
	arg1(c, d)
	a, b := 10, 100
	arg3(&a, &b)
	fmt.Println("a =", a, "b =", b)
	var e int = 1
	e = return1(a, b)
	fmt.Println(e)
}

func arg1(a int, b int) {
	a = a + b
	// return
	// fmt.Println("不会被输出")
}

func arg2(a, b int) {
	a = a + b
}

func arg3(a, b *int) {
	*a = *a + *b
	*b = 888
}

func no_arg() {
	fmt.Println("gogogo")
}

func return1(a, b int) int {
	a += b
	c := a
	return c
}

func return2(a, b int) (c int) {
	a += b
	c = a
	return c
}

func return3() (int, int) {
	now := time.Now()
	return now.Hour(), now.Minute()
}
