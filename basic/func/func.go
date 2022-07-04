package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//带余数的除法
func div1(a, b int) (int, int) {
	return a / b, a % b
}

//可以在定义函数返回值类型的同时定义返回值的变量
func div2(a, b int) (q int, r int) {
	q = a / b
	r = a % b
	return q, r
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func eval1(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		q, _ := div1(a, b) //如果不想用第二个返回值，通过_接收想要忽略的返回值
		return q
	default:
		panic("Unsupport Operators")

	}
}

//一般函数返回值会包含一个error类型的变量，用来告知调用者函数执行是否发生错误
func eval2(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div1(a, b) //如果不想用第二个返回值，通过_接收想要忽略的返回值
		return q, nil
	default:
		return 0, fmt.Errorf("Error: Unsupported Operators")

	}
}

//回调函数,函数本身也可以作为函数的参数
func apply(op func(a, b int) int, a, b int) int {
	//获取op的函数名称
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("Calling function: ", opName)
	return op(a, b)
}

//可变参数列表,numbers可以当作数组
func sum(numbers ...int) (result int) {
	for _, x := range numbers {
		result += x
	}
	return result

}
func main() {
	fmt.Println(div1(19, 5))
	fmt.Println(div2(19, 5))

	fmt.Println(eval1(19, 5, "/"))
	// fmt.Println(eval1(19, 5, "ws"))

	fmt.Println(eval2(19, 5, "/"))
	fmt.Println(eval2(19, 5, "ws"))

	fmt.Println(apply(pow, 2, 10))

	//匿名函数
	fmt.Println(apply(func(a, b int) int {
		return a + b
	}, 2, 10))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
