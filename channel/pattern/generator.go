package main

import (
	"fmt"
	"math/rand"
	"time"
)

//字符串生成器
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("Message from service %s: %d", name, i)
			i++
		}
	}()
	return c
}
func main() {
	m := msgGen("m") //m、n可以理解为服务
	n := msgGen("n") //
	for {
		fmt.Println(<-m)
		fmt.Println(<-n)
	}
}
