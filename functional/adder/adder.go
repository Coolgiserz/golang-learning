package main

import "fmt"

//"传统"的函数式编程不能有变量/“状态”，因此不能像下面那样在函数中声明sum:=0；go语言的函数式编程允许有变量
func adder() func(int) int {
	sum := 0

	return func(i int) int {
		sum += i
		return sum
	}
}
func main() {
	a := adder()
	for i := 0; i < 10; i += 1 {
		fmt.Println(a(i))
	}
}
