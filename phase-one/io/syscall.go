package io

import (
	"bytes"
	"fmt"
	"os/exec"
)

// 执行系统命令
func SysCall() {
	// 查看系统命令所在目录，确保命令已安装
	cmd_path, err := exec.LookPath("go")
	if err != nil {
		fmt.Println("could not found command go")
	}
	fmt.Printf("command go in path %s\n", cmd_path)

	cmd := exec.Command("go", "version") // 相当于go version
	// cmd.Output()运行命令并获得其输出结果
	if output, err := cmd.Output(); err != nil {
		fmt.Println("got output failed", err)
	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("python", "C:\\Users\\18101\\Desktop\\go\\go_learning\\phase-one\\io\\hello.py")

	// cmd.Output()运行命令并获得其输出结果
	if output, err := cmd.Output(); err != nil {
		fmt.Println("python execute failed", err)
	} else {
		fmt.Println(string(output))
	}

	cmd = exec.Command("rm", "../data/biz.log") // err := os.Remove("../data/biz.log")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	// 如果不需要获得命令的输出，直接调用cmd.Run()即可
	err = cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String()) // 错误信息
	} else {
		fmt.Println(out.String())
	}
}
