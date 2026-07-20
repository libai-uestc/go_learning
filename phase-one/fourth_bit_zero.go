package main

import (
	"fmt"
)

// IsFourthBitZero 判断一个 uint 的第4位（从右往左数，即索引3）是否为 0
func IsFourthBitZero(n uint) bool {
	// 1 << 3 会生成一个只有第4位为1，其他位全为0的掩码 (二进制: 0000 1000)
	// 将 n 与该掩码进行“按位与”操作：
	// 如果 n 的第4位是 0，结果就是 0
	// 如果 n 的第4位是 1，结果就大于 0
	return (n & (1 << 3)) == 0
}

func main14() {
	// 测试用例 1：8 的二进制是 1000，它的第 4 位（索引3）是 1
	var num1 uint = 8
	fmt.Printf("数字 %d (二进制: %04b) 的第4位是否为 0? %t\n", num1, num1, IsFourthBitZero(num1))

	// 测试用例 2：7 的二进制是 0111，它的第 4 位（索引3）是 0
	var num2 uint = 7
	fmt.Printf("数字 %d (二进制: %04b) 的第4位是否为 0? %t\n", num2, num2, IsFourthBitZero(num2))

	// 测试用例 3：15 的二进制是 1111，它的第 4 位（索引3）是 1
	var num3 uint = 15
	fmt.Printf("数字 %d (二进制: %04b) 的第4位是否为 0? %t\n", num3, num3, IsFourthBitZero(num3))
}
