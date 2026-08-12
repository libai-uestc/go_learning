package main

import (
	"fmt"
	projectprepare "libai/go/basic/phase-one/project_prepare"
)

func InitLogger() {
	fmt.Println("init logger")
	fmt.Println("main是否匹配正则表达式", projectprepare.Reg.Match([]byte("hello123")))
}

func main() {
	projectprepare.CheckReg()
	InitLogger()
	InitDatabase()
}

func InitDatabase() {
	fmt.Println("init database")
}
