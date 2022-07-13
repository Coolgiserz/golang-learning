package main

import (
	"fmt"
	"sync"
)

/**
channel.go中为了保证所有goroutine均得到执行，使用了time.Sleep(time.Millisecond)的代码强制主协程进行休息，这种做法并不好
本文件演示通过Channel等待任务结束的方法和使用系统工具使主协程适时（等子协程均执行完）退出

1. 通过Channel等待任务结束的方法：子协程中启用新协程，新建一个通道通知子协程中任务是否完成，主协程接收来自新协程的chan中的数据

**/
type worker struct {
	channel chan int
	done    chan bool
}

//=========bufferedChannel2: 实际上还是串行=============
func bufferedChannel2() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)

	}
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'a'
		<-workers[i].done
	}
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'A'
		<-workers[i].done
	}
}
func createWorker(id int) worker {
	w := worker{
		channel: make(chan int),
		done:    make(chan bool),
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w worker) {
	for {
		fmt.Printf("Worker-----(id, message): (%d, %c)\n", id, <-w.channel)
		go func() {
			w.done <- true
		}()
	}
}

//=============bufferedChannel3:  并行，但可以进一步封装、优化========

func bufferedChannel3() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i)

	}
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'A'
	}
	for i := 0; i < 10; i++ {
		<-workers[i].done
	}

	for i := 0; i < 10; i++ {
		<-workers[i].done
	}
}
func createWorker3(id int) worker {
	w := worker{
		channel: make(chan int),
		done:    make(chan bool),
	}
	go doWorker3(id, w)
	return w
}

func doWorker3(id int, w worker) {
	for {
		fmt.Printf("Worker-----(id, message): (%d, %c)\n", id, <-w.channel)
		go func() {
			w.done <- true //这里可以进一步封装
		}()
	}
}

//=====bufferedChannel4
type worker4 struct {
	channel chan int
	done    func()
}

//======wait group=======
func bufferedChannel4() {
	var workers [10]worker4
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		workers[i] = createWorker4(i, &wg)

	}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'a'
	}
	for i := 0; i < 10; i++ {
		workers[i].channel <- i + 'A'
	}
	wg.Wait() // 一直阻塞到WaitGroup 计数器为0

}
func createWorker4(id int, wg *sync.WaitGroup) worker4 {
	w := worker4{
		channel: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker4(id, w)
	return w
}

func doWorker4(id int, w worker4) {
	for {
		fmt.Printf("Worker-----(id, message): (%d, %c)\n", id, <-w.channel)
		w.done() //通知WaitGroup计数器，当前子协程已完成->计数器-1
	}
}
func main() {
	bufferedChannel4()
}
