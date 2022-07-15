package main

import (
	"fmt"
	"math/rand"
	"time"
)

//字符串生成器
// func msgGen(name string) chan string {
// 	c := make(chan string)
// 	go func() {
// 		i := 0
// 		for {
// 			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
// 			c <- fmt.Sprintf("Message from service %s: %d", name, i)
// 			i++
// 		}
// 	}()
// 	return c
// }
//【问题：如果希望协程接收外界指令/消息（如是否中断程序）；可以引入done chan】
func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Duration(rand.Intn(5000)) * time.Millisecond):
				c <- fmt.Sprintf("Message from service %s: %d", name, i)
			case <-done: //接收来自main函数的指令，done通道收到数据则中断此协程
				fmt.Println("Interruct")
				//确保main函数在该子协程退出后再退出
				done <- struct{}{}
				return

			}

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

//非阻塞等待
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default: //一旦m:=<-c阻塞则会进入default分支
		return "", false
	}
}

//超时机制
func timeoutWait(c chan string, duration time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(duration): //一旦m:=<-c阻塞则会进入default分支
		return "", false
	}
}

/*什么时候用普通fanIn什么时候用fanInBySelect？
看是否明确知道工作的goroutine的数目，不明确可以用fanIn，加上可变参数；明确知道参数数量用select，更加简洁
*/
func main() {
	done := make(chan struct{})
	// m := msgGen("1", done) //m、n可以理解为服务
	n := msgGen("2", done) //
	// n1 := msgGen("3") //

	// result := fanIn(m, n)
	// result := fanInBySelect(m, n)
	// result := fanInUncertained(m, n)
	// result := fanInUncertained(m, n, n1)

	// for {

	// fmt.Println(<-m)
	// fmt.Println(<-n)
	// fmt.Println(<-result)

	//========测试非阻塞等待=========
	// fmt.Println("====非阻塞等待nonBlockingWait====")
	// if nReceived, ok := nonBlockingWait(n); ok {
	// 	fmt.Println(nReceived)
	// } else {
	// 	fmt.Println("No data")
	// }

	// fmt.Println("====超时机制timeoutWait====")
	// if nReceived, ok := timeoutWait(n, 2*time.Second); ok {
	// 	fmt.Println(nReceived)
	// } else {
	// 	fmt.Println("Timeout!")
	// }

	// }

	//测试中断协程
	for i := 0; i < 5; i++ {
		fmt.Println("====超时机制timeoutWait====")
		if nReceived, ok := timeoutWait(n, 2*time.Second); ok {
			fmt.Println(nReceived)
		} else {
			fmt.Println("Timeout!")
		}
	}
	done <- struct{}{}
	<-done //确保子协程退出后main才退出
	// time.Sleep(time.Second)
}
