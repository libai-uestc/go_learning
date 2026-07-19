package main

import "fmt"

func main4() {
	var pi float32 = 3.14
	const pi_1 float32 = 3.1415926
	pi = 3.26
	_ = pi

	const (
		E  = 2.71
		PI = 3.14
	)

	var (
		var_1 = 3.11
		var_2 = 3.12
	)

	_ = var_1
	_ = var_2

	const (
		APPLE  = 1
		HUAWEI = 0
		VIVO   = 0.5
	)

	const (
		a = iota // 0
		b
		c
		d
		_
		e = 7
		f        // 7
		g = iota // 7
		h        // 8
	)
	fmt.Println(h)

	const (
		NOT_USE = 1 << (10 * iota) // iota = 0
		KB      = 1 << (10 * iota)
		MB      = 1 << (10 * iota)
		GB      = 1 << (10 * iota)
	)

	const (
		NOT_USE_1 = 1 << (10 * iota)
		KB_1
		MB_1
		GB_1
	)

	const (
		m = 20
		n // 20
		p // 20
		q // 20
	)

	const (
		ss, mm = iota + 1, iota + 2 // 1,2 iota = 0
		cc, dd
		ff, nn
	)
}
