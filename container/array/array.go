package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 999
	fmt.Println("数组", arr)
}

func printArray1(arr *[5]int) {
	arr[0] = 999
	fmt.Println("数组", arr)
}
func main() {

	//声明数组
	var arr1 [5]int
	arr2 := [3]int{4, 3, 2}
	arr3 := [...]int{4, 2, 2, 1, 6} //编译器自动推断长度。切片类型
	var grid [4][5]int              //二维数组：4行5列
	fmt.Println(arr1, arr2, arr3, grid)

	//遍历数组
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	for i := 0; i < len(arr3); i += 1 {
		fmt.Println(i, arr3[i])
	}

	//数组类型：值类型，作为参数传递给函数时会进行拷贝
	printArray(arr3)
	printArray(arr1)
	fmt.Println(arr3, arr1) //printArray中改变数组的值不影响原来的arr1、arr3的值
	// printArray(arr2) //会报错，[3]int和[5]int是不同的类型

	printArray1(&arr3)
	printArray1(&arr1)
	fmt.Println(arr3, arr1)

	fmt.Println(arr1[2:5])
}
