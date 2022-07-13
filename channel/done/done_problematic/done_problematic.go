package main

import (
	"fmt"
	"time"
)

/**
试执行如下代码，可以发现程序退出时打印出来的Worker可能不足20个，这因为部分子协程还未执行主协程便已经退出。
在bufferedChannel2的最后一句加上time.Sleep(time.Millisecond)语句可以使子协程均有足够时间完成，
但这种做法属于硬编码，不够优雅、扩展性不强。

**/

//=========带缓冲的通道=============
func bufferedChannel2() {
	var channels [10]chan<- int //第二个参数为缓冲区大小/容量，可用于提升性能
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		channels[i] <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		channels[i] <- i + 'A'
	}
	time.Sleep(time.Millisecond)
}
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker-----(id, message): (%d, %c)\n", id, <-c)
	}
}

// func doWorker(id int, c chan<- int) {
// 	c = make(c)
// 	for n := range c {

// 		fmt.Printf("(id, message): (%d, %c)\n", id, n)

// 	}
// }
func main() {
	// testChan5()
	bufferedChannel2()
}
