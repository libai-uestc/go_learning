package io

import "os"

const (
	logText = "君不见黄河之水天上来，奔流到海不复回。君不见高堂明镜悲白发，朝如青丝暮成雪。\n"
)

// 带缓冲的FileWriter
//
// Note: 不支持并发。golang自带的bufio.NewWriter支持并发
type BufferedFileWriter struct {
	buffer         []byte   // 缓存的内容
	bufferEndIndex int      // buffer里有效内容的结束位置
	fout           *os.File // 文件句柄
}

func NewWriter(fout *os.File, bufferSize int) *BufferedFileWriter {
	return &BufferedFileWriter{
		buffer:         make([]byte, bufferSize), // len=cap=bufferSize
		bufferEndIndex: 0,
		fout:           fout,
	}
}

// 向文件中写入内容。（大概率只是写入了缓存，还没有真正写入磁盘）
func (w *BufferedFileWriter) Write(cont []byte) {
	if len(cont) >= len(w.buffer) { // 要写的内容比缓存空间还要大，则直接把cont写到文件里去
		w.Flush()
		w.fout.Write(cont)
	} else {
		// 先预判buffer能否容下cont
		if w.bufferEndIndex+len(cont) > len(w.buffer) { // 不能容下
			w.Flush()
		}
		// append2(w.buffer,w.bufferEndIndex,cont)
		copy(w.buffer[w.bufferEndIndex:], cont) // golang内置的copy函数功能上等价于自己写的append2函数，但比append2函数更高效
		w.bufferEndIndex += len(cont)
	}
}

// 把buffer里的内容全部写入磁盘文件
func (w *BufferedFileWriter) Flush() {
	w.fout.Write(w.buffer[0:w.bufferEndIndex]) // 把buffer里的内容写入文件
	w.bufferEndIndex = 0                       //清空buffer
}

// 把src拷贝到dest[index:]里去
func append2(dest []byte, index int, src []byte) {
	for i := 0; i < len(src); i++ {
		dest[index+i] = src[i]
	}
}
