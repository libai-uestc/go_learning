package main

import (
	"fmt"
	"io"
	myhttp "libai/go/basic/phase-one/web"
	"net/http"
	"os"
	"strings"
)

func HttpObservation() {
	fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
	resp, err := http.Get("http://127.0.0.1:5678/obs?name=libai&age=18")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("response proto: %s\n", resp.Proto) // http1.1
	if major, minor, ok := http.ParseHTTPVersion(resp.Proto); ok {
		fmt.Printf("http major version %d, http minor version %d\n", major, minor)
	}
	fmt.Printf("response status: %s\n", resp.Status)          // 200 OK
	fmt.Printf("response status code: %d\n", resp.StatusCode) // 200

	for key, values := range resp.Header {
		fmt.Printf("%s:%v\n", key, values)
		if key == "Date" {
			if tm, err := http.ParseTime(values[0]); err == nil {
				fmt.Printf("server time %s\n", tm.Format("2006-01-02 15:04:05"))
			}
		}
	}

	fmt.Println("response body:")
	io.Copy(os.Stdout, resp.Body) // 两个io数据流的拷贝
	os.Stdout.WriteString("\n\n")

}

func Get() {
	fmt.Println(strings.Repeat("*", 30) + "GET" + strings.Repeat("*", 30))
	resp, err := http.Get("http://127.0.0.1:5678/get?" + myhttp.EncodeUrlParams(map[string]string{"name": "白 Li", "age": "18"}))
	if err != nil {
		panic(err)
	} else {
		defer resp.Body.Close()
		fmt.Printf("response status: %s\n", resp.Status)
		fmt.Println("response body:")
		// io.Copy(os.Stdout, resp.Body) // 两个io数据流的拷贝
		if body, err := io.ReadAll(resp.Body); err == nil {
			fmt.Print(string(body))
		}
		os.Stdout.WriteString("\n\n")
	}
}

func main() {
	HttpObservation()
	Get()
}
