package main

import (
	"fmt"
	//两个模块都实现了Retriever接口，实现接口不需要显式声明实现了哪个接口
	"coolgiserz.com/learngo/interfaces/retrievers/mock"
	"coolgiserz.com/learngo/interfaces/retrievers/real"
)

type Retriever interface {
	Get(url string) string
	// Test() string
}

type Poster interface {
	Post(url string)
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {

}
func main() {
	var r Retriever
	r = mock.Retriver{"Fake"}
	fmt.Printf("%T %v \n", r, r)
	// fmt.Println(download(r))
	var rq Retriever
	rq = &real.Retriever{UserAgent: "Haha"}
	// fmt.Println(download(rq))
	fmt.Printf("%T %v \n", rq, rq)
}
