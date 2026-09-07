package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"libai/go/basic/phase-one/web/live/demo/util"
	"log"
	"net/http"
	"os"
	"strconv"
	// "time"
)

func GetFileSize(FileName string) int {
	resp, err := http.Get("http://127.0.0.1:5678/file_size/" + FileName)
	if err != nil {
		log.Printf("GetFileSize failed: %s", err)
		return -1
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(os.Stdout, resp.Body)
		return -1
	}

	bs, _ := io.ReadAll(resp.Body)
	FileSize, err := strconv.Atoi(string(bs))
	if err != nil {
		log.Printf("返回的文件大小不是纯数字: %s", string(bs))
		return -1
	}
	return FileSize
}

func Download(FileName string) {
	FileSize := GetFileSize(FileName)
	if FileSize <= 0 {
		return
	}
	log.Printf("文件总大小是: %d B\n", FileSize)
	file, err := os.OpenFile("demo/client/"+FileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) // 当你第一次打开这个文件，如果之前这个文件它不存在的话，就临时创建这个文件，如果之前已经存在过的，就清空掉，重新下载，打开文件是为了读写文件
	if err != nil {
		log.Printf("open file failed: %s", err)
		return
	}
	defer file.Close()

	resp, err := http.Get("http://127.0.0.1:5678/download/" + FileName)
	if err != nil {
		log.Printf("download file failed: %s", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		io.Copy(os.Stdout, resp.Body)
		return
	}

	// 从响应头里读出文件哈希值
	hash := resp.Header.Get("hash")

	bar := util.Bar{}
	bar.NewOptionWithGraph(0, int64(FileSize), "#") // 初始化进度条
	acc := 0                                        // 累计下载的字节数
	buffer := make([]byte, 2048)
	for {
		n, err := resp.Body.Read(buffer)
		if err != nil {
			if err == io.EOF {
				file.Write(buffer[:n])
				acc += n
				bar.Play(int64(acc)) // 刷新进度条
			} else {
				log.Printf("read response body failed: %s", err)
			}
			break
		}
		file.Write(buffer[:n])
		acc += n
		bar.Play(int64(acc)) // 刷新进度条
		// time.Sleep(10 * time.Millisecond)
	}
	fmt.Println()

	myHash := Hash4File("demo/client/" + FileName)
	if myHash != hash {
		log.Printf("文件下载数据有丢失，%s!=%s", myHash, hash)
	} else {
		log.Println("文件下载完全没问题", myHash)
	}
}

func main() {
	Download("Michael_Jordan’s_Top_60_Career _Plays.mp4")
}

func Hash4File(FileName string) string {
	file, err := os.Open(FileName) // 打开文件使用Open，打开文件写入的话，使用OpenFile
	if err != nil {
		return ""
	}
	defer file.Close()

	hasher := sha256.New()

	buffer := make([]byte, 2048)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF { // End Of File
				hasher.Write(buffer[:n]) // 响应体
			} else {
				log.Printf("读文件异常: %s", err)
			}
			break
		}
		hasher.Write(buffer[:n]) // 响应体
	}

	return hex.EncodeToString(hasher.Sum(nil))
}

// certutil -hashfile .\demo\client\通天.mp4 SHA256
// 980f27b29450e9627c4580ce97d22c46587ab59a32f027e071dc1b0ab98b136f

//  go run .\demo\client\
// 2026/09/07 18:56:03 文件总大小是: 321060911 B
// [##################################################]100% 321060911/321060911
// 2026/09/07 18:56:07 文件下载完全没问题 5a88f2911e369be4abe6b6379400e0ea707372e7ea621dea7580fcd8ba28e167
