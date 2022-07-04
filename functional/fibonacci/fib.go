package main

import (
	"bufio"
	"fmt"
	"io"

	// "io"
	"os"
	"strings"
)

type intGen func() int

//类似生成器
func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fibWriter := bufio.NewWriter(file)
	defer fibWriter.Flush()
	f := fib()
	for i := 0; i < 30; i += 1 {
		fmt.Fprintln(fibWriter, f())
	}
}

//go语言也能给函数实现接口，方法的“接收者”也是参数，函数式编程的理念中，函数可作为参数，因此也可以作为接收者
// func printFileContents(reader io.Reader) {
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

//为函数类型intGen实现io.Reader接口
//type Reader interface {
// 	Read(p []byte) (n int, err error)
// }
//
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1024 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d \n", next)
	return strings.NewReader(s).Read(p)
}
func main() {
	// f := fib()
	// printFileContents(f) //f是intGen类型，已经实现了io.Reader接口，因此可以作为printFileContents函数的参数
	// for i := 0; i < 10; i += 1 {
	// 	fmt.Println(f())
	// }
	writeFile("fib.txt")
}
