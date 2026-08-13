package io

import (
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	if fin, err := os.Open("../data/verse.txt"); err != nil {
		fmt.Printf("open file failed: %v\n", err) // 比如文件不存在
	} else {
		defer fin.Close()
		bs := make([]byte, 100)
		fin.Read(bs)
		fmt.Println(string(bs))

		fin.Seek(0, 0)
		fin.Read(bs)
		fmt.Println(string(bs))

		fin.Seek(0, 0)
		const BATCH = 10
		buffer := make([]byte, BATCH)
		for {
			n, err := fin.Read(buffer)
			if n > 0 {
				fmt.Println(buffer[0:n])
			}
			if err == io.EOF {
				break
			}
		}
	}
}
