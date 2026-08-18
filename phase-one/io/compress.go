package io

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

const (
	_ = iota
	GZIP
	ZLIB
)

func Compress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	stat, _ := fin.Stat()
	fmt.Printf("压缩前文件大小 %dB\n", stat.Size())

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	var writer io.WriteCloser
	switch compressAlgo {
	case GZIP:
		writer = gzip.NewWriter(fout)
	case ZLIB:
		writer = zlib.NewWriter(fout)
	}

	io.Copy(writer, fin) // io.Copy会开辟一个32K的buffer，分批拷贝数据。如果想自定义buffer大小，可以使用io.CopyBuffer
	writer.Close()
	fout.Close()
	fin.Close()
}

// 解压
func Decompress(inFile, outFile string, compressAlgo int) {
	fin, err := os.Open(inFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := fin.Stat()
	fmt.Printf("压缩后文件大小 %dB\n", stat.Size())

	var reader io.ReadCloser
	switch compressAlgo {
	case GZIP:
		reader, _ = gzip.NewReader(fin)
	case ZLIB:
		reader, _ = zlib.NewReader(fin)
	}

	fout, err := os.OpenFile(outFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(fout, reader)
	reader.Close()
	fin.Close()
	fout.Close()
}
