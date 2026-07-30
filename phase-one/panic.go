package main

import "fmt"

/*
何时会发生panic：
1. index out of range和divede by zero是常见的runtime error，会发生panic
2. 通常系统初始化发生重大问题会主动调用panic(any)

panic时会依次执行：
1. 执行已经注册的defer(后注册的先执行)，未注册的defer不执行
2. 打印错误调用堆栈
3. 调用os.Exit(2)结束go进程
4. 如果recover()函数执行了，第二步和第三步就不执行了
*/

func defer_panic() {
	// panic(43)
	defer fmt.Println(1)
	var arr []int
	n := 0
	// defer fmt.Println(1/n) //在注册defer时就要计算1/n，发生panic
	defer func() {
		_ = arr[n]
		_ = 1 / n //defer func 内部发生panic，main协程不会exit，其他defer还可以正常执行
		defer fmt.Println(2)
	}()
	defer fmt.Println(3)
}

func soo() {
	fmt.Println("enter soo")

	// defer func() {
	// 	// recover 必须在defer中才能生效
	// 	if panicInfo := recover(); panicInfo != nil {
	// 		fmt.Printf("soo函数中发生了panic:%v\n", panicInfo) // 任何类型都可以通过%v来进行格式化输出
	// 		// debug.PrintStack() // 打印调用堆栈
	// 	}
	// }()

	fmt.Println("regist recover")

	defer fmt.Println("hello")
	defer func() {
		n := 0
		_ = 3 / n
		defer fmt.Println("how are you!")
	}()

	defer func() {
		// recover 必须在defer中才能生效
		if panicInfo := recover(); panicInfo != nil {
			fmt.Printf("soo函数中发生了panic:%v\n", panicInfo) // 任何类型都可以通过%v来进行格式化输出
			// debug.PrintStack() // 打印调用堆栈
		}
	}()
}

func main() {
	// defer_panic()
	soo()
}
