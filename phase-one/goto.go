package main

import "fmt"

func basic_goto() {
	var i int = 4
MY_LABEL:
	i += 3 //锚点，Lable
	i *= 2
	fmt.Println(i)
	if i > 200 {
		return
	}
	goto MY_LABEL
}

func if_goto() {
	var i int = 4
	if i%2 == 0 {
		goto L1 // Label指示的是某一个代码，并没有圈定一个代码块，所以goto L1会把25到29行的代码全部执行
	} else {
		goto L2
	}
L1:
	i += 3
	fmt.Println(i)
L2:
	i *= 3
	fmt.Println(i)
}

func for_goto() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)

	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					goto L1
				} else if j%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func continue_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					continue L1
				} else if i%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func break_label() {
	const SIZE = 5
L1:
	for i := 0; i < SIZE; i++ {
	L2:
		fmt.Printf("开始检查第%d行\n", i)
	L3:
		if i%2 == 1 {
			for j := 0; j < SIZE; j++ {
				fmt.Printf("开始检查第%d列\n", j)
				if j%3 == 0 {
					break L1
				} else if i%3 == 1 {
					goto L2
				} else {
					goto L3
				}
			}
		}
	}
}

func main10() {
	// basic_goto()
	// if_goto()
	// for_goto()
	// continue_label()
	break_label()
}
