package main

import (
	"fmt"
	"time"
)

func ticker() {
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tk := time.NewTicker(time.Second)
	// defer tk.Stop()
	for i := 0; i < 10; i++ {
		<-tk.C
		fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	}
	tk.Stop()
}

func timer() {
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(time.Second)
	<-tm.C
	// <-tm.C
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm.Stop()

	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	<-time.After(time.Second)
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
}

func main26() {
	// ticker()
	timer()
}
