package main

import "fmt"

type Age int // Age和int可以互相做强制类型转换。存储的数据类型（成员变量）是一样的，但行为（成员方法）是不一样

type Tall = int // Tall和int完全等价，不需要显式做类型转换

type Odt Shiper

func (Age) move(src string, dest string) (int, error) {
	return 0, nil
}

func (Age) whistle(n int) int {
	return 0
}

func main31() {
	var a int
	var b Age
	var c Tall

	c = a + Tall(10)
	b = Age(a) + 10 // 字面量不需要显示显式转为Age类型

	fmt.Println(a, b, c)

	var o Odt
	o.tonage = 100
	o.name = "libai"
	seaTransport("BJ", "SH", Shiper(o))
	transport("BJ", "SH", b)
}
