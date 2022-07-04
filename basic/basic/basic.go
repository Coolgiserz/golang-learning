package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa      = 5 //不能使用:=
	bb int  = 6
	cc bool = true
)

//声明变量，未初始化变量时，变量值默认取0和空字符串，而非像C++那样发生不明行为
func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

// Go 语言也支持自动类型推断
func variableInitValue() {
	var a, b int = 1, 2
	var s string = "Hello"
	fmt.Printf("%d %d %q\n", a, b, s)
}

func variableInitValueTypeInference() {
	var a, b, c = 1, 2, false
	var s = "Hello type inference"
	fmt.Println(a, b, c, s)
}

func variableShorterInit() {
	a, b := 1, 2 // :=只能在函数内使用，在包内函数外无法使用
	s := "Hello shorter initialize variable"
	fmt.Println(a, b, s)
}

func euler() {
	c := cmplx.Exp(1i*math.Pi) + 1
	// fmt.Println("欧拉公式： ", c, cmplx.Abs(c))
	fmt.Printf("欧拉公式： %.3f\n", c)

}

func forceTypeTransform() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b))) //前面声明了c为int类型，则需要对math.Sqrt返回的float64类型进行强制转换为int才能赋值给c
	fmt.Println("强制类型转换", c)
}

//常量
func consts() {
	const name = "Coolgiser"
	const x, y float64 = 3, 4 //const如果不指定类型，数值可以当作各种类型来用
	var c int
	c = int(math.Sqrt((x*x + y*y)))
	fmt.Println("常量：", c)
}

//枚举
func enumerate() {
	//普通枚举
	const (
		gnn int = 0
		cnn int = 1
	)
	fmt.Println("普通枚举: ", gnn, cnn)
	//自增枚举
	const (
		b = 1 << (10 * iota)
		b1
		b2
		b3
	)
	fmt.Println("iota: ", b, b1, b2, b3)

}
func main() {
	fmt.Println("Golang Basic:")
	variableZeroValue()
	variableInitValue()
	variableInitValueTypeInference()
	variableShorterInit()
	fmt.Println("全局变量（可以在包内任何函数中使用）: ", aa, bb, cc)
	euler()
	forceTypeTransform()
	consts()
	enumerate()
}
