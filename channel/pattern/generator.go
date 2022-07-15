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

//从两个通道中等待数据，谁快就先收谁
func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-c1
		}
	}()

	go func() {
		for {
			c <- <-c2
		}
	}()
	return c
}

//从不确定数量的通道中等待数据，谁快就先收谁
func fanInUncertained(chs ...chan string) chan string {
	c := make(chan string)

	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch) //ch通过值传递，保证内部每一个for循环中都有独一份当前ch的拷贝
	}

	return c
}

//同时等待多个服务的返回
func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()

	return c
}

/*什么时候用普通fanIn什么时候用fanInBySelect？
看是否明确知道工作的goroutine的数目，不明确可以用fanIn，加上可变参数；明确知道参数数量用select，更加简洁
*/
func main() {
	m := msgGen("1")  //m、n可以理解为服务
	n := msgGen("2")  //
	n1 := msgGen("3") //

	// result := fanIn(m, n)
	// result := fanInBySelect(m, n)
	// result := fanInUncertained(m, n)
	result := fanInUncertained(m, n, n1)

	for {

		// fmt.Println(<-m)
		// fmt.Println(<-n)
		fmt.Println(<-result)
	}
}
