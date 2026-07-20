package main

import "fmt"

func Set_Fourth_Bit_To_Zero(n uint) uint {
	return n &^ (1 << 3) // 1 << 3 = 1000 ,^1000 = 0111  ， n&(0111)->第四位置为0
}

func main15() {
	// 测试用例 1：15 的二进制是 1111。清除第 4 位后，应该变成 0111 (即 7)
	var num1 uint = 15
	result1 := Set_Fourth_Bit_To_Zero(num1)
	fmt.Printf("原数字: %2d (二进制: %04b)\n", num1, num1)
	fmt.Printf("修改后: %2d (二进制: %04b)\n\n", result1, result1)

	// 测试用例 2：8 的二进制是 1000。清除第 4 位后，应该变成 0000 (即 0)
	var num2 uint = 8
	result2 := Set_Fourth_Bit_To_Zero(num2)
	fmt.Printf("原数字: %2d (二进制: %04b)\n", num2, num2)
	fmt.Printf("修改后: %2d (二进制: %04b)\n\n", result2, result2)

	// 测试用例 3：7 的二进制是 0111。它的第 4 位本来就是 0，修改后应该保持不变，仍为 0111 (即 7)
	var num3 uint = 7
	result3 := Set_Fourth_Bit_To_Zero(num3)
	fmt.Printf("原数字: %2d (二进制: %04b)\n", num3, num3)
	fmt.Printf("修改后: %2d (二进制: %04b)\n", result3, result3)
}
