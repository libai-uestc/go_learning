package main

import "fmt"

func main40() {
	var i int
	fmt.Printf("i=%v\n", i)
	var f float32
	fmt.Printf("f=%v\n", f)
	var b byte
	fmt.Printf("b=%v\n", b)
	var bl bool
	fmt.Printf("bl=%v\n", bl)
	var s string
	fmt.Printf("s=%v\n", s)
	var p *int
	fmt.Printf("p=%v\n", p)
	type User struct {
		Gender bool
		Age    int
	}
	var u User
	fmt.Printf("u=%v\n", u)
	var err error
	fmt.Printf("err=%v\n", err)
	err = nil
	var arr [3]int
	fmt.Printf("arr=%v\n", arr)
	var slc []string
	fmt.Printf("slc=%v\n", slc)
	var mp map[int]bool
	fmt.Printf("mp=%v\n", mp)
	var ch chan int
	fmt.Printf("ch=%v\n", ch)
	var ifc any
	v1, ok := ifc.(float64)
	fmt.Println(v1, ok)
	ifc = 3.14
	v2, ok := ifc.(float32)
	fmt.Println(v2, ok)
	v3, ok := ifc.(float32)
	fmt.Println(v3, ok)

	mapValue, exists := mp[7]
	fmt.Printf("mapValue=%v\n", mapValue)
	_ = exists

	ch = make(chan int, 10)
	close(ch)
	chValue, ok := <-ch
	fmt.Printf("chValue=%v\n", chValue)
	_ = ok

	// err:=gorm.Select("*").Where("age=0").First(&u).Error() 查不到结果时u也是0值。通过err来告诉你是不是0值
}
