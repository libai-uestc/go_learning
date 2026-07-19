package main

import "fmt"

// 结构体：把多个变量进行打包构成了一个整体
type User struct {
	// 成员变量
	Id            int
	Score         float64
	Name, address string
}

// 成员方法
func (user User) hello() {
	fmt.Println("My Name is", user.Name)
}

func main5() {
	var u User
	u = User{Id: 10, Score: 100, Name: "libai", address: "China"} //Score可以不赋值，默认是0值，整数或浮点数不赋值默认是0，布尔默认是false，字符串默认是空字符串
	fmt.Println(u.Name)
	u = User{99, 100, "libai", "China"}
	u.hello()

	// 匿名结构体 只在一个场景，不需要重新使用
	type student struct {
		Name string
		Age  int
	}

	var student_1 struct {
		Name string
		Age  int
	}

	student_1.Name = "libai"
	student_1.Age = 18
	u.Name = student_1.Name
	a := student_1.Age
	_ = a

	u2 := User{}
	u2.address = "sh"

	u3 := u2
	fmt.Println(u3.Name)

	u3_1 := &u2 //取址符号，u3是*User类型
	fmt.Println(u3_1.Name)

	u4 := new(User) //new是先创建空结构体，再返回其指针  u4是*User类型
	u4.Name = "libai_u4"
	fmt.Println(u4.Name)
}
