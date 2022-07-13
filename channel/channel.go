package main

import (
	"fmt"
	"time"
)

//！会发生死锁，因为chan中的数据没人接收
func testChan1() {
	c := make(chan int) //新建一个int类型的通道
	c <- 1
	fmt.Println(c)
}

// 此时不会死锁，有新的协程接收数据
func testChan2() {
	c := make(chan int) //新建一个int类型的通道,chan也是一等公民，可以作为参数和返回值
	go func() {
		n := <-c //此时c为匿名函数外部的参数
		fmt.Println(n)
	}()
	c <- 1
}

func worker(id int, channel chan int) {
	for {
		fmt.Printf("Worker %d received channel %c \n", id, <-channel) //%c 打印字符，%d打印数字
	}
}

// 批量创建channel
func testChan3() {
	var channels [10]chan int

	for i := 0; i < 10; i++ {
		channels[i] = make(chan int) //批量创建channel
		go worker(i, channels[i])
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Second) //如果不加这行，则可能程序会提前退出，部分协程没机会执行导致无法接收channel的数据与打印消息
}

// chan可以作为返回值
func createWorker(id int) chan int { //此时不能确定返回的chan是用于接收数据还是发送数据，可以进一步明确chan的行为，见createWorker1
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received channel %c \n", id, <-c) //%c 打印字符，%d打印数字
		}

	}()
	return c
}
func testChan4() {
	fmt.Println("=======TestChan4========")
	var channels [10]chan int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Second) //如果不加这行，则可能程序会提前退出，部分协程没机会执行导致无法接收channel的数据与打印消息
}

func createWorker1(id int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received channel %c \n", id, <-c) //%c 打印字符，%d打印数字
		}

	}()
	return c
}

// func testChan5() {
// 	fmt.Println("=======TestChan5========")
// 	var channels [10]<-chan int

// 	for i := 0; i < 10; i++ {
// 		channels[i] = createWorker1(i)
// 	}

// 	for i := 0; i < 10; i++ {
// 		channels[i] <- 'a' + i //会报错invalid operation: channels[i] <- 'a' + i (send to receive-only type <-chan int)

// 	}
// 	time.Sleep(time.Second) //如果不加这行，则可能程序会提前退出，部分协程没机会执行导致无法接收channel的数据与打印消息
// }

func createWorker2(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received channel %c \n", id, <-c) //%c 打印字符，%d打印数字
		}

	}()
	return c
}

func testChan6() {
	fmt.Println("=======TestChan6========")
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker2(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //不报错。因为chan<-类型可以从goroutine中接收数据
	}
	time.Sleep(time.Second) //如果不加这行，则可能程序会提前退出，部分协程没机会执行导致无法接收channel的数据与打印消息
}

/**
	chan<- vs <-chan
	chan<-表示只能chan智能接收来自另一goroutine的数据
	<-chan表示只能从chan中接收数据

**/

//=========带缓冲的通道=============
func bufferedChannel1() {
	bc := make(chan int, 3) //第二个参数为缓冲区大小/容量，可用于提升性能
	go worker1(0, bc)
	bc <- 'a'
	bc <- 'b'
	bc <- 'c'
	bc <- 'e'
	bc <- 'f'
	time.Sleep(time.Millisecond)
}

func worker1(id int, c chan int) {
	for {
		fmt.Printf("(id, message): (%d, %c)\n", id, <-c)
	}
}
func bufferedChannel2() {
	bc := make(chan int, 3)       //第二个参数为缓冲区大小/容量，可用于提升性能
	go worker2_judgeClose2(0, bc) //需要判断chan是否被关闭，否则如果发送方关闭了chan，worker2会循环打印接收到的类型的0值
	bc <- 'a'
	bc <- 'b'
	bc <- 'c'
	bc <- 'e'
	bc <- 'f'
	close(bc) //由发送方close
	time.Sleep(time.Millisecond)
}

//需要判断chan是否被关闭，否则如果发送方关闭了chan，worker2会循环打印接收到的类型的0值
func worker2_judgeClose1(id int, c chan int) {
	for {
		if n, ok := <-c; ok {
			fmt.Printf("(id, message): (%d, %c)\n", id, n)
		}

	}
}

func worker2_judgeClose2(id int, c chan int) {
	for n := range c {

		fmt.Printf("(id, message): (%d, %c)\n", id, n)

	}
}
func main() {
	// testChan5()
	bufferedChannel2()
}
