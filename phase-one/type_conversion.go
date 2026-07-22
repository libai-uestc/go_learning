package main

import (
	"fmt"
	"math"
	"strconv"
)

func string_other_convert() {
	var err error
	var i int = 8
	var i64 int64 = int64(i)
	// int -> string
	var s string = strconv.Itoa(i) // 内部调用FormatInt,Integer -> a(string，字符串)
	s = strconv.FormatInt(i64, 10)
	// string -> int
	i, err = strconv.Atoi(s)
	// string -> int64
	i64, err = strconv.ParseInt(s, 10, 64)

	// float -> string
	var f float64 = 8.123456789
	s = strconv.FormatFloat(f, 'f', 2, 64) // 保留2位小数 %2f
	fmt.Println(s)
	// string -> float
	f, err = strconv.ParseFloat(s, 64)

	// string -> []byte
	var arr []byte = []byte(s)
	// []byte -> string
	s = string(arr)

	// string -> []rune
	var brr []rune = []rune(s)
	// []rune -> string
	s = string(brr)

	fmt.Printf("err %v\n", err)
}

func cast() {
	// 高精度->低精度，数字很小的时候这种转换没问题
	var ua uint64 = 1
	i8 := int8(ua)
	fmt.Printf("i8=%d\n", i8)

	// 最高位的1变成了符号位
	ua = uint64(math.MaxUint64)
	i64 := int64(ua)
	fmt.Printf("i64=%d\n", i64) // -1。负数用补码形式存储

	// 位数丢失
	ui32 := uint32(ua) //64位转为32位，会有丢失
	fmt.Printf("ui32=%d\n", ui32)

	// 单个字符可以转为int
	var i int = int('中') //'中'是字符，对于字符而言，实际上就是一个rune，就是UTF-8编码集，总能对应到一个编号
	fmt.Printf("i=%d\n", i)

	// bool和int不能相互转换

	// byte和int可以相互转换
	var by byte = byte(i)
	i = int(by)
	fmt.Printf("i=%d\n", i)

	// float和int可以互相转换，小数位会丢失
	var ft float32 = float32(i)
	i = int(ft)
	fmt.Printf("i=%d\n", i)
}

func main23() {
	var test_int int = 100
	var test_char string = "李"
	var test_rune rune = '白'
	var i_to_string string = strconv.Itoa(test_int)
	var err error
	i_to_string = strconv.FormatInt(int64(test_int), 10)
	test_int, err = strconv.Atoi(test_char)
	if err != nil {
		fmt.Println("转换失败")
	}
	var int_64 int64 = 100
	int_64, err = strconv.ParseInt(test_char, 10, 64)
	var num_str string = "2024"
	test_int, err = strconv.Atoi(num_str)
	if err == nil {
		fmt.Println("转换成功！")
	}
	rune_to_string := string(test_rune)
	fmt.Printf("%c -> %q", test_rune, rune_to_string)
	fmt.Printf("final: %s%s", test_char, rune_to_string)
	_ = i_to_string
	_ = int_64
}
