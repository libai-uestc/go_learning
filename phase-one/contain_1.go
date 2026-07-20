package main

func CountOnes(n uint64) int {
	count := 0
	for n > 0 {
		n &= (n - 1)
		count++
	}
	return count
}

// 29 = 11101
// 11101 & 11100 = 11100
// 11100 & 11011 = 11000
// 11000 & 10111 = 10000
// 10000 & 01111 = 00000
// count = 4
// Brian Kernighan 算法
