package main

import (
	"fmt"
	"unicode/utf8"
)

//go如何支持多语言？
//如何编码中文、ascii码？
func main() {
	s := "你好Go"            //UTF-8编码：英文1字节、中文3字节
	fmt.Println(s, len(s)) //len获得的是字节长度

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %c) ", i, ch)
	}

	fmt.Println()
	fmt.Println(utf8.RuneCountInString(s)) //获得字符数量

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c \n", ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
