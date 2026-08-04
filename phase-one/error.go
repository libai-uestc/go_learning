package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found error") // error变量名一般以Err开头
)

// 自定义error
type MyError struct {
	Name string
	Code int
	Desc string
}

// 构造函数
func NewMyError(name string, code int, desc string) *MyError {
	return &MyError{
		Name: name,
		Code: code,
		Desc: desc,
	}

	// err := new(NewMyError)
	// err.Name = name
	// err.Code = code
	// err.Desc = desc
	// return *err
}

// 拥有Error() string即实现了error接口
func (e MyError) Error() string {
	return fmt.Sprintf("[%d]%s %s", e.Code, e.Name, e.Desc)
}

// Print对象时默认会调用对象的String()方法。error是个例外
func (e MyError) String() string {
	return e.Name
}

// 如果一个函数返回多个值，error一般为最后一个值
func divide(a, b int) (int, error) {
	if b == 0 {
		// return 0, ErrNotFound
		// return 0, fmt.Errorf("divide error %d %d", a, b)
		return 0, NewMyError("math", 101, "divide by zero")
	} else {
		return a / b, nil
	}
}

// 调函数后应该优先判断error是否为nil

func main36() {
	err2 := NewMyError("math", 101, "divide by zero")
	fmt.Printf("%s\n", err2)
	c, err := divide(5, 0)
	if err != nil {
		fmt.Printf("出错 %s\n", err)
	} else {
		fmt.Printf("结果 %d\n", c)
	}
}
