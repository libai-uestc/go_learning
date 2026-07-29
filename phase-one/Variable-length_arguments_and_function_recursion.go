package main

import "fmt"

func variable_length_arg(a int, other ...int) int {
	sum := a
	for _, ele := range other {
		sum += ele
	}
	if len(other) > 0 {
		fmt.Printf("first ele %d len %d cap %d\n", other[0], len(other), cap(other))
	} else {
		fmt.Printf("len %d cap %d\n", len(other), cap(other))
	}
	return sum
}

func sum(arr ...int) int {
	s := 0
	if len(arr) == 0 {
		return s
	}
	s += arr[0]
	s += sum(arr[1:]...)
	return s
}
