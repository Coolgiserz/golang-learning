/**
	（应用场景、解决的痛点）
	场景：需要从多个通道收指令/数据，但不确定哪个通道的数据来得更快。
	问题：一般来说通道上收发数据是会阻塞的，一个通道在其数据还没被接收时会阻塞
	使用select进行调度
	同时从多个通道中接收数据，

**/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func testSelect1() {
	fmt.Println("====testSelect1====")
	var c1, c2 chan int //此时两个chan int都是nil，nil chan不会被select

	for {
		select { //多路复用、非阻塞式通道；c1、c2哪个通道先输出数据，n就接收谁的数据，如果c1、c2出数据的时间相同，则随机选择
		case n := <-c1:
			fmt.Println("received from c1: ", n)
		case n := <-c2:
			fmt.Println("received from c2: ", n)
		default: //加上default就编程非阻塞式地从chan获取值
			fmt.Println("No data")
		}
	}
}

func testSelect2() {
	fmt.Println("====testSelect2====")
	c1, c2 := chanGenerator(), chanGenerator() //创建非空chan

	for {
		select { //多路复用、非阻塞式通道；c1、c2哪个通道先输出数据，n就接收谁的数据，如果c1、c2出数据的时间相同，则随机选择
		case n := <-c1:
			fmt.Println("received from c1: ", n)
		case n := <-c2:
			fmt.Println("received from c2: ", n)
			// default:
			// 	fmt.Println("No data")
		}
	}
}
func chanGenerator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

//==========testSelect3============
//select中接收到来自某个通道的数据后如何使用
func testSelect3() {
	fmt.Println("====testSelect3====")
	c1, c2 := chanGenerator(), chanGenerator() //nil chan
	w := createWorkerChan(0)
	hasValue := false
	n := 0
	for {
		var activeWorker chan<- int //activeWorker默认为nil，无法被select

		if hasValue { //如果从被监听通道中接收到值，则启用activeWorker，保证其是非空chan
			activeWorker = w
		}
		select { //多路复用、非阻塞式通道；c1、c2哪个通道先输出数据，n就接收谁的数据，如果c1、c2出数据的时间相同，则随机选择
		case n = <-c1:
			hasValue = true
			// fmt.Println("received from c1: ", n)
		case n = <-c2:
			hasValue = true
			// fmt.Println("received from c2: ", n)
		case activeWorker <- n: //【问题：如果activeWorker消耗n的速度过慢，按目前的实现部分n会丢失，还需要考虑将n值进行存储，确保都被worker利用，见testSelect4】
			hasValue = false
			fmt.Printf("{n}=%d has been sent to Worker {w}\n", n)

		}
	}
}

type worker struct {
	channel chan int
	done    chan bool
}

func createWorkerChan(id int) chan<- int {
	w := make(chan int)
	go doWorker(id, w)
	return w
}

func doWorker(id int, w chan int) {
	for n := range w {
		time.Sleep(1 * time.Second)
		fmt.Printf("+++++++Worker %d--received--(message): (%d)\n", id, n)

	}
}

//==========testSelect4============
//
//<演示计时器使用>
//1. testSelect3程序无法自动退出，可以通过引入计时器使得程序在经过一定时间后结束；
//2. 如果想实现 一定时间内未收到数据打印timeout，可以在select中添加一个case，用计时器计算两次收发数据时间差
//
// <如何查看积压了多少value？(收到的数据比发送出去的数据快，会导致数据积压)>
// 通过time.Tick（每隔一定时间返回一个值）定时报告积压值
//
//=========================
func testSelect4() {
	fmt.Println("====testSelect4====")
	c1, c2 := chanGenerator(), chanGenerator() //nil chan
	w := createWorkerChan(0)
	n := 0
	tick := time.Tick(2 * time.Second) //定时返回一个chan time.Time
	timer := time.After(10 * time.Second)
	var values []int //引入缓存
	for {
		var activeWorker chan<- int //activeWorker默认为nil，无法被select

		var value int
		if len(values) > 0 { //如果从被监听通道中接收到值，则启用activeWorker，保证其是非空chan
			activeWorker = w
			value = values[0]
		}
		select { //多路复用、非阻塞式通道；c1、c2哪个通道先输出数据，n就接收谁的数据，如果c1、c2出数据的时间相同，则随机选择
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- value: //接收到值
			fmt.Printf("{n}=%d has been sent to Worker {w}\n", n)
			values = values[1:]
		case <-time.After(600 * time.Millisecond):
			fmt.Printf("timeout!!!!!!!\n")
		case <-tick: //定时打印values的大小
			fmt.Println("len(values)=", len(values))
		case <-timer: //到时间终止程序
			fmt.Println("BYEBYE!!")
			return
		}
	}
}
func main() {
	// testSelect1()
	// testSelect2()
	// testSelect3()
	testSelect4()

}
