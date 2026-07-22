package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func asign_string() {
	s1 := "My name is 李白"

	s2 := "He say:\"I'm fine.\" \n\\Thank\tyou. \\"

	s3 := `here is first line.
	
	there is third line.
	`
	fmt.Println("s1")
	fmt.Println(s1)
	fmt.Println("s2")
	fmt.Println(s2)
	fmt.Println("s3")
	fmt.Println(s3)
}

func string_impl() {
	s1 := "My name is 李白"
	arr := []byte(s1)
	brr := []rune(s1)
	fmt.Printf("last byte %d\n", arr[len(arr)-1]) // string可以转换为[]byte或[]rune类型
	fmt.Printf("last byte %c\n", arr[len(arr)-1]) // %c以unicode字符串进行输出
	fmt.Printf("last rune %d\n", brr[len(brr)-1])
	fmt.Printf("last rune %c\n", brr[len(brr)-1])
	L := len(s1)
	fmt.Printf("string len %d byte array len %d rune array len %d\n", L, len(arr), len(brr))
	for _, ele := range s1 {
		fmt.Printf("%c ", ele)
	}
	fmt.Println()
	for i := 0; i < L; i++ {
		fmt.Printf("%c ", s1[i]) // 这里用的是byte切片
	}
	fmt.Println()

	arr[0] = 9
	// s1[0]=9 //字符串不能修改
	fmt.Println(utf8.RuneCountInString(s1), len([]rune(s1))) //查看string里有几个rune

}

func string_join() {
	s1 := "Hello"
	s2 := "how"
	s3 := "are"
	s4 := "you"
	merged := s1 + " " + s2 + " " + s3 + " " + s4
	fmt.Println(merged)
	merged = fmt.Sprintf("%s %s %s %s", s1, s2, s3, s4)
	fmt.Println(merged)
	merged = strings.Join([]string{s1, s2, s3, s4}, " ")
	fmt.Println(merged)
	sb := strings.Builder{}
	sb.WriteString(s1)
	sb.WriteString(" ")
	sb.WriteString(s2)
	sb.WriteString(" ")
	sb.WriteString(s3)
	sb.WriteString(" ")
	sb.WriteString(s4)
	sb.WriteString(" ")
	merged = sb.String()
	fmt.Println(merged)
}

func string_op() {
	s := "born to win, born to die."
	fmt.Printf("sentence length %d\n", len(s))
	fmt.Printf("\"s\" length %d\n", len("s")) //英文字母的长度为1
	fmt.Printf("\"中\" length %d\n", len("中")) //一个汉字占3个长度
	arr := strings.Split(s, " ")
	fmt.Printf("arr[3]=%s\n", arr[3])
	fmt.Printf("contain die %t\n", strings.Contains(s, "die"))
	fmt.Printf("contain wine %t\n", strings.Contains(s, "wine"))
	fmt.Printf("first index of born %d\n", strings.Index(s, "born"))
	fmt.Printf("last index of born %d\n", strings.LastIndex(s, "born"))
	fmt.Printf("begin with born %t\n", strings.HasPrefix(s, "born")) //前缀
	fmt.Printf("end with die %t\n", strings.HasSuffix(s, "die."))    //后缀
}

func main() {
	asign_string()
	string_impl()
	string_join()
	string_op()
}
