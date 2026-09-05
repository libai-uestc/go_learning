package main

import (
	"fmt"
	"html/template"
	"io"
	myhttp "libai/go/basic/phase-one/web"
	"net/http"
	"strconv"
	"strings"
	"time"
	// "os"
)

func HttpObservation(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request method: %s\n", r.Method)
	fmt.Printf("request host: %s\n", r.Host) // 服务端host
	fmt.Printf("request url: %s\n", r.URL)
	fmt.Printf("request proto: %s\n", r.Proto)
	fmt.Println("request header")
	for key, values := range r.Header { // 变量类型是map,map的key是字符串，map的value是字符串切片
		fmt.Printf("%s: %v\n", key, values)
	}
	fmt.Println()
	fmt.Printf("request body: ")
	// io.Copy(os.Stdout,r.Body) // 把r.Body流里的内容拷贝到os.Stdout流里
	if body, err := io.ReadAll(r.Body); err == nil {
		fmt.Println(string(body))
	}
	fmt.Println()

	w.Header().Add("tRAce-id", "4723956498105") // Trace-Id
	w.WriteHeader(http.StatusBadRequest)        // 如果这一行被注释掉，则会默认把状态码设置为http.StatusOK
	w.Write([]byte("Hello Boy\n"))
	w.Write([]byte("Hello Girl\n"))
	fmt.Fprint(w, "Hello Boy\n") // 和w.Write([]byte("Hello Boy\n"))的效果是一样的
	fmt.Println(strings.Repeat("*", 60))
}

func Get(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	params := myhttp.ParseUrlParams(r.URL.RawQuery)
	fmt.Fprintf(w, "your name is %s, age is %s\n", params["name"], params["age"])
	fmt.Println(strings.Repeat("*", 60))
}

// 流式传输海量数据
func HugeBody(w http.ResponseWriter, r *http.Request) {
	line := []byte("Heavy is the head who wears the crown.\n")
	const R = 10 // line重复发送几次
	totalSize := R * len(line)
	w.Header().Add("content-length", strconv.Itoa(totalSize))
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "不支持flush", http.StatusInternalServerError)
		return
	}
	for i := 0; i < R; i++ {
		if _, err := w.Write(line); err != nil { // 即使不显式Flush()，Write()的内容足够多(大几K)时也会触发Flush()
			fmt.Printf("%d send error: %s\n", i, err)
			break
		} else {
			flusher.Flush() // 强制write to tcp
			time.Sleep(time.Second)
		}
	}
	fmt.Println(strings.Repeat("*", 60))
}

func Student(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生产模版对象
	tmpl, err := template.ParseFiles("./phase-one/web/server/student.tmpl") // 相对于执行go run的路径
	if err != nil {
		fmt.Println("create template failed:", err)
		return
	}
	type Student struct {
		Id     int
		Name   string
		Gender string
		Score  int
	}
	// 利用给定数据渲染模板，并将结果写入w
	students := []Student{{1, "张三", "男", 80}, {2, "李四", "女", 40}, {3, "王五", "女", 50}}
	tmpl.Execute(w, students)
}

func main() {
	// 定义路由
	http.HandleFunc("/obs", HttpObservation)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/stream", HugeBody)
	http.HandleFunc("/student", Student)
	// 启动Http Server
	if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
		panic(err)
	}
}
