package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//简单数列求和
func sumFromN1ToN2(n1 int, n2 int) int {
	var result int
	for i := n1; i <= n2; i += 1 {
		result += i
	}
	return result
}

//将十进制数转为二进制
func convert2bin(n int) string {
	var result string
	for ; n > 0; n /= 2 {
		tmp := n % 2
		result = strconv.Itoa(tmp) + result //strconv.Itoa(tmp)把int类型变量转换成string
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("error: %s", err))
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { //相当于遍历迭代器，
			fmt.Println(scanner.Text())
		}
	}
}

//死循环
func foreverloop() {
	for {
		println("forever")
	}
}
func main() {
	fmt.Println(sumFromN1ToN2(1, 100))
	fmt.Println(convert2bin(13))
	fmt.Println(convert2bin(0))
	fmt.Println(convert2bin(4))
	printFile("data/test.txt")
	foreverloop()
}
