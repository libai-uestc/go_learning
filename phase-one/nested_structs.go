package main

import "fmt"

type Residence struct {
	Province string
	City     string
}

type User struct {
	Id            int
	Score         float64
	Name, address string
	residence     Residence
	info          struct { //匿名结构体
		Age   int
		mouth int
		year  int
	}
	Father *User // 指针的初始化就是nil
}

type TreeNode struct {
	Data       int
	LeftChild  *TreeNode
	RightChild *TreeNode
}

func (user User) hello() {
	fmt.Println("My Name is", user.Name)
}

func main6() {
	var user User
	user.Score = 100
	user = User{Id: 32, address: "China", Name: "libai"}
	fmt.Println(user.Name)
	u := User{12, 100, "gogogo", "China_sh", Residence{"bj", "yz"},
		struct {
			Age   int
			mouth int
			year  int
		}{18, 12, 2026},
		&User{},
	}
	fmt.Printf("%v\n", u)
	fmt.Printf("%+v\n", u)
	fmt.Printf("%#v\n", u)
	u.hello()

	ue := u
	ue.Name = "lb"
	_ = ue
}
