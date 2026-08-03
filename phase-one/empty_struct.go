package main

import (
	"fmt"
	"reflect"
	"time"
	"unsafe"
)

type ETS struct {
}

// 所有的空结构体指向同一个地址（内核是完全一样的）
func allEmptyStructIsSame() {
	var a ETS
	var b ETS
	var c struct{} // 空结构体类型，类似的有interface{}
	fmt.Printf("address of a %p b %p c %p\n", &a, &b, &c)
	fmt.Printf("size of a %d b %d c %d\n", unsafe.Sizeof(a), unsafe.Sizeof(b), unsafe.Sizeof(c))
	fmt.Printf("size of a %d b %d c %d\n", reflect.TypeOf(a).Size(), reflect.TypeOf(b).Size(), reflect.TypeOf(b).Size()) // 反射
}

// 空结构体的应用场景
func scenariosOfEmptyStruct() {
	// 使用map间接实现set
	set := map[int]struct{}{
		1: struct{}{},
		3: struct{}{},
		5: struct{}{},
	}
	if _, exists := set[5]; exists {
		fmt.Println("5是存在的")
	} else {
		fmt.Println("5是不存在的")
	}

	// 可以通过WaitGroup等待
	blocker := make(chan struct{}, 0)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("done")
		blocker <- struct{}{}
	}()
	<-blocker // 等待子协程结束
}

func main35() {
	allEmptyStructIsSame()
}
