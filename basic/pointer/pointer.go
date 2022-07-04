package main

import "fmt"

func swap1(a, b int) {
	b, a = a, b
}

func swap2(a, b *int) {
	*b, *a = *a, *b
}

func swap3(a, b int) (int, int) {
	return b, a
}
func main() {
	a, b := 1, 2
	// swap1(a, b)
	// fmt.Println(a, b)

	// swap2(&a, &b) //把a、b变量的地址按值传递给swap2函数
	// fmt.Println(a, b)

	a, b = swap3(a, b)
	fmt.Println(a, b)
}
