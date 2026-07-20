package main

import "fmt"

func main13() {
	var a uint64 = 5
	fmt.Printf("a=%b\n", a)
	binaryFormat(a)
}

func binaryFormat(a uint64) {
	var c uint64 = 1 << 63
	for i := 0; i < 64; i++ {
		if c&a == c {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
		c = c >> 1
	}
}
