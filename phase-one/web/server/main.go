package main

import (
	"encoding/json"
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

func Post(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if ct, exists := r.Header["Content-Type"]; exists {
		switch ct[0] {
		case "text/plain":
			io.Copy(w, r.Body) // 直接把请求体作为响应体。io.Copy内部会去根据情况调用Sendfile、Splice等零拷贝策略
		case "application/json":
			body, err := io.ReadAll(r.Body)
			if err == nil {
				params := make(map[string]string, 10)
				if err := json.Unmarshal(body, &params); err == nil {
					fmt.Fprintf(w, "your name is %s, age is %s\n", params["name"], params["age"])
				}
			} else {
				fmt.Println("read request body error", err)
			}
		case "application/x-www-form-unlencoded":
			body, err := io.ReadAll(r.Body)
			if err == nil {
				fmt.Println("request body", string(body))
				params := myhttp.ParseUrlParams(string(body))
				fmt.Fprintf(w, "your name is %s, age is %s\n", params["name"], params["age"])
			} else {
				fmt.Println("read request body error", err)
			}
		}
	}
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

func Cookie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request header:")
	for key, value := range r.Header {
		fmt.Println(key, value)
	}
	// 其实可以直接通过r.Cookies()获得*http.Cookie，没必要自己解析
	if values, exists := r.Header["Cookie"]; exists {
		cookies, _ := http.ParseCookie(values[0]) // 多个request cookie全在values[0]里
		fmt.Println("request cookie:")
		for _, cookie := range cookies {
			fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
		}
		fmt.Println(strings.Repeat("*", 60))
	}

	// Set-Cookie
	expiration := time.Now().Add(30 * 24 * time.Hour)
	cookie1 := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration, Domain: "localhost", Path: "/"}
	cookie2 := http.Cookie{Name: "jwt", Value: "1234", Expires: expiration, Domain: "localhost", Path: "/"}
	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)
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

func router1() {
	// 路由
	http.HandleFunc("/obs", HttpObservation)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/post", Post)
	http.HandleFunc("/stream", HugeBody)
	http.HandleFunc("/cookie", Cookie)
	http.HandleFunc("/student", Student)

	// 启动Http Server
	if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
		panic(err)
	}
}

func router2() {
	if err := http.ListenAndServe("127.0.0.1:5678", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && r.URL.Path == "/obs" {
			HttpObservation(w, r)
		} else if r.Method == http.MethodGet && r.URL.Path == "/get" {
			Get(w, r)
		} else if r.Method == http.MethodPost && r.URL.Path == "/post" {
			Post(w, r)
		} else if r.Method == http.MethodGet && r.URL.Path == "/stream" {
			HugeBody(w, r)
		} else if r.Method == http.MethodGet && r.URL.Path == "/cookie" {
			Cookie(w, r)
		} else if r.Method == http.MethodGet && r.URL.Path == "/student" {
			Student(w, r)
		}
	})); err != nil {
		panic(err)
	}

}

func router3() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /obs", func(w http.ResponseWriter, r *http.Request) {
		HttpObservation(w, r)
	})
	mux.HandleFunc("GET /get", func(w http.ResponseWriter, r *http.Request) {
		Get(w, r)
	})
	mux.HandleFunc("POST /post", func(w http.ResponseWriter, r *http.Request) {
		Post(w, r)
	})
	mux.HandleFunc("GET /stream", func(w http.ResponseWriter, r *http.Request) {
		HugeBody(w, r)
	})
	mux.HandleFunc("GET /cookie", func(w http.ResponseWriter, r *http.Request) {
		Cookie(w, r)
	})
	// restful风格参数
	mux.HandleFunc("GET /get/{name}/{age}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "your name is %s, age is %s\n", r.PathValue("name"), r.PathValue("age"))
	})

	mux.HandleFunc("GET /student", func(w http.ResponseWriter, r *http.Request) {
		Student(w, r)
	})

	if err := http.ListenAndServe("127.0.0.1:5678", mux); err != nil {
		panic(err)
	}
}

func main() {
	// // 定义路由
	// http.HandleFunc("/obs", HttpObservation)
	// http.HandleFunc("/get", Get)
	// http.HandleFunc("/stream", HugeBody)
	// http.HandleFunc("/student", Student)
	// http.HandleFunc("/post", Post)
	// http.HandleFunc("/cookie", Cookie)
	// // 启动Http Server
	// if err := http.ListenAndServe("127.0.0.1:5678", nil); err != nil {
	// 	panic(err)
	// }
	// router1()
	// router2()
	router3()
}
