package main

import "fmt"

func main2() {
	var TestDefaultValue int      // 变量声明
	fmt.Println(TestDefaultValue) // 使用变量
	var a int = 8
	var b = a
	_ = b
	c := b // :声明变量 var
	a = c

	var (
		d uint16
		e int8
		f float32
		g float64
	)

	a = -5
	d = 05
	a = 0o57
	a = 0xab3
	a = 1_2_3_4
	a = 100_000_000 //100M
	f = 1.432
	g = 34.
	m := 12. // 默认float64类型
	var n bool = true
	_, _, _, _ = d, e, f, m

	fmt.Printf("a=%d, g=%f, n=%t\n", a, g, n)
	v := 1.12342351236
	fmt.Printf("v=%g,v=%e", v, v) // %g默认6位小数，%e科学计数法 .2f的话就只有小数点后两位
	u := 123
	i := 456
	fmt.Printf("u=%[1]d, i=%[2]d,i=%[2]d, u=%[1]d\n", u, i)
}
