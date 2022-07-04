package main

import (
	"fmt"
	"io/ioutil"
)

func branchIf1() {
	filename := "data/test.txt"

	if text, err := ioutil.ReadFile(filename); err == nil { // if的条件前可以接赋值语句
		fmt.Printf("%s \n", text)

	} else {
		// fmt.Println("Error occurred!", err)
		panic(err) //程序主动调用panic，立即返回函数
	}
}

func branchIf2(variable int) {

	if variable > 50 {
		fmt.Println("variable > 50")
	} else if variable > 10 {
		fmt.Println("variable > 10")
	} else {
		fmt.Println("variable <= 10")
	}
}

//case后不需要接break，默认匹配到一个case执行完即会退出switch，不像c++还会往下匹配其它case
func branchSwitch1(variable int) {
	switch {
	case variable > 50:
		fmt.Println("variable>50")
	case variable > 10:
		fmt.Println("variable > 10")
	default:
		fmt.Println("variable <= 10")
	}
}

// 使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func branchSwitch2(variable int) {
	switch {
	case variable > 50:
		fmt.Println("variable > 50")
		fallthrough
	case variable < 50:
		fmt.Println("fallthrough variable < 50")
		// fallthrough
	default:
		fmt.Println("variable <= 10")
	}
}

// Type-Switch判断某个interface变量存储的变量类型
func branchSwitch3() {
	var variable interface{}
	switch variable.(type) {
	case int:
		fmt.Println("type: int")

	case float32:
		fmt.Println("type: float32")

	case nil:
		fmt.Println("type: nil")

	default:
		fmt.Println("other type")

	}
}

func main() {
	branchIf1()
	branchIf2(100)
	branchIf2(40)
	branchIf2(10)
	branchSwitch1(12)
	branchSwitch2(950)
	branchSwitch3()
}
