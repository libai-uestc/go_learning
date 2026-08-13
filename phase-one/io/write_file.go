package io

import (
	"fmt"
	"os"
)

func WriteFile() {
	// 如果使用go test，则相对路径是相对于xxx_test.go文件的路径
	// 如果使用go run或编译后直接运行，则相对路径是相对于执行命令时所在的路径
	if fout, err := os.OpenFile("../data/verse.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o666); err != nil {
		fmt.Printf("open file failed: %s\n", err.Error())
	} else {
		defer fout.Close()
		fout.WriteString("李白\n")
		fout.WriteString("千金散尽还复来")
		fout.WriteString("\n")
		fout.Write([]byte("天生我材必有用"))
		fout.WriteString("\n")
	}
}
