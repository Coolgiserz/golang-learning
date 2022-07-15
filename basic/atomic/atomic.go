package main

import (
	"fmt"
	"sync"
	"time"
)

// type atomicInt int //原子；atomic在传统并发编程中意思是线程安全的
type atomicInt struct {
	value int
	lock  sync.Mutex
} //原子；atomic在传统并发编程中意思是线程安全的

func (a *atomicInt) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

//可以为go run 添加-race参数检测数据访问冲突
func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
