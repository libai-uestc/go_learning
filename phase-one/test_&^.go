package main

import "fmt"

func main16() {
	var a int = 10     // 1010
	b := a &^ (1 << 3) // 1000        0111 & 1010 =0010
	fmt.Println(b)

}
