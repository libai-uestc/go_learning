package main

import (
	"fmt"
	"strings"
)

func rangeChannel() {
	var ch chan int
	if ch == nil {
		fmt.Printf("ch is nil,ch len %d cap %d\n", len(ch), cap(ch))
	}
	// ch <- 2 // 不能向nil chan里发送数据
	if len(ch) == 0 {
		fmt.Println("ch length is 0")
	}
	ch = make(chan int, 8)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Printf("ch len %d cap %d\n", len(ch), cap(ch))
	v := <-ch
	fmt.Println(v)
	v = <-ch
	fmt.Println(v)
	fmt.Println()

	close(ch) //关闭掉意味着不能写入新元素了
	// 遍历并取走（receive）管道里的元素。当管道里已无剩余元素且没有close管道时，receive操作会一直阻塞，最终报deadlock。当管道为空且被close后，for循环推出
	for ele := range ch {
		fmt.Println(ele)
	}

	c := make(chan int, 10)
	send(c)
	recv(c)
}

// 这个函数表示只能写入到这个chan，出了这个函数就没有限制了，只限制当前函数
func send(c chan<- int) {
	c <- 1
}

func recv(c <-chan int) {
	v := <-c
	fmt.Printf("take %d from read-only channel\n", v)
}

func changeArray(arr [3]int) { //有数组或...表示数组
	arr[0]++
}

func changeSlice(slc []int) {
	if len(slc) > 0 {
		slc[0]++
	}
}

func changeMap(mp map[int]bool) {
	mp[0] = true
}

func changeChan(ch chan bool) {
	if cap(ch) > len(ch) {
		ch <- true
	}
}

func main() {
	rangeChannel()
	fmt.Println(strings.Repeat("*", 50))

	arr := [3]int{}
	changeArray(arr)
	fmt.Println(arr[0]) //还是0

	slc := []int{1, 2, 3}
	changeSlice(slc)
	fmt.Println(slc[0]) //修改为2了

	mp := map[int]bool{0: false}
	changeMap(mp)
	fmt.Println(mp[0])

	ch := make(chan bool, 3)
	changeChan(ch)
	fmt.Println(<-ch)
}
