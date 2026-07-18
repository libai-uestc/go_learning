package main

import "fmt"

// test 注释
func main() {
	fmt.Printf("A=%d Z=%d\n", 'A', 'Z')
	var base int = 'Z' - 'A' + 1 //进制
	fmt.Println(base, "进制")      //26
	//总和
	var total int
	total += 'D' - 'A' + 1
	total += base * ('F' - 'A' + 1)
	total += base * base * ('X' - 'A' + 1)
	fmt.Println("total", total) //16384
}
