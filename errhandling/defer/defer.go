package main

import "fmt"

func tryDefer() {

	defer fmt.Println("Hello")  //使得该行代码在函数结束时执行
	defer fmt.Println("Golang") // defer 后进先出 ，因此函数结束时会先输出Golang再输出Hello
	fmt.Println("Hahaha")
	panic("sss")
}
func main() {
	tryDefer()
}
