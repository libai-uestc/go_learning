package main

import "fmt"

func main() {
	type User struct {
		Name string
		Age  int
	}
	type Vedio struct {
		Length int
		Name   string
		Author User
	}
	type Vedio_1 struct {
		Length int
		Name   string
		User   //匿名成员
	}
	u := User{Name: "libai", Age: 19}
	v := Vedio{Length: 100, Name: "gogogo", Author: u}
	fmt.Println(v.Length)
	fmt.Println(v.Name)
	fmt.Println(v.Author.Name)
	fmt.Println(v.Author.Age)
	w := Vedio_1{Length: 100, Name: "gogogogogo", User: u}
	fmt.Println(w.User.Name) //访问父类的name
	fmt.Println(w.Name)      // 访问自己的name

}
