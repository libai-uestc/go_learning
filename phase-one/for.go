package main

import "fmt"

func main9() {
	var sum int
	for a, b := 9, 30; a >= 0 && b > a; a, b = a-1, b/2 {
		sum += a
		if a == 4 {
			break
		}
		if a == 5 {
			continue
		}
	}
	fmt.Println(sum)
}
