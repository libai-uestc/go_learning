package main

import (
	"fmt"
	"io"
	myhttp "libai/go/basic/phase-one/web"
	"net/http"
	"strings"
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

	for i := 0; i < R; i++ {
		if _, err := w.Write(line); err != nil { // 即使不显式Flush()，Write()的内容足够多(大几K)时也会触发Flush()
			fmt.Printf("%d send error: %s\n", i, err)
			break
		}
	}
	fmt.Println(strings.Repeat("*", 60))
}

func main() {
	// 定义路由
	http.HandleFunc("/obs", HttpObservation)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/stream", HugeBody)

	// 启动Http Server
	if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
		panic(err)
	}
}
