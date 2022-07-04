package main

import "fmt"

//函数参数为[]int时，表明参数类型是int型的切片
func updateSlice(slice []int) {
	slice[0] = 999
}

func printSlice(slice []int) {
	fmt.Printf("slice: %v, len: %d, cap: %d \n", slice, len(slice), cap(slice))
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5}
	s := arr[1:3] //半开半闭区间，与python一致
	fmt.Println("arr[1:3]=", s)
	fmt.Println("arr[:3]=", arr[:3])
	fmt.Println("arr[1:]=", arr[1:])
	fmt.Println("arr[:]=", arr[:])

	//对slice的修改会改变底层的array
	fmt.Println("Before Update Slice s")
	fmt.Println(s)
	fmt.Println("After Update Slice s")
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)

	//切片扩展：slice可以向后扩展不可向前扩展
	//每个slice底层包含ptr、len、cap
	//向后扩展不可超越数组容量
	// fmt.Println(arr[6]) // out of bound error
	fmt.Printf("%v, %d, %d \n", s, len(s), cap(s))

	//通过make函数创建指定长度和容量的切片
	s3 := make([]int, 5)
	s4 := make([]int, 10, 16)
	printSlice(s3)
	printSlice(s4)

	//删除切片元素,如删除第3个元素
	fmt.Println("===删除切片元素===")
	s5 := []int{4, 2, 1, 0}
	printSlice(s5)
	s5 = append(s5[:2], s5[3:]...)
	printSlice(s5)
}
