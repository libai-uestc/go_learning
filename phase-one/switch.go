package main

import (
	"fmt"
	"strings"
)

func switch_basic() {
	color := "yellow"
	switch color {
	case "green":
		fmt.Println("go")
	default:
		fmt.Println("stop")
	}

	color = "green"
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("stop")
	default:
		fmt.Println("all not!")
	}

	color = "black"
	switch color {
	case "green":
		fmt.Println("go")
	case "red":
		fmt.Println("stop")
	case "yellow":
		fmt.Println("停！")
	default:
		fmt.Printf("非法的信号灯！🚥:%s\n", strings.ToUpper(color))
	}
}

func add(a int) int {
	return a + 10
}

func switch_expression() {
	var a int = 5
	switch add(a) {
	case 15:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}

func switch_condition() {
	color := "yellow"
	switch color {
	case "green":
		fmt.Println("go")
	case "red", "yellow":
		fmt.Println("stop")
	}
	switch {
	case add(5) > 10:
		fmt.Println("right")
	default:
		fmt.Println("wrong")
	}
}

func fall_through(age int) {
	fmt.Printf("您的年龄是%d，您可以：\n", age)
	switch {
	case age > 50:
		fmt.Println("当总统")
		fallthrough
	case age > 25:
		fmt.Println("结婚")
		fallthrough
	case age > 22:
		fmt.Println("study")
		fallthrough
	case age > 18:
		fmt.Println("drive")
		fallthrough
	case age > 3:
		fmt.Println("study study study!!!")
	}
}

func main() {
	switch_basic()
	switch_expression()
	switch_condition()
	fall_through(20)
}
