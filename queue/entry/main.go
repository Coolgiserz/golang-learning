package main

import (
	"fmt"

	"coolgiserz.com/learngo/queue"
)

func main() {
	q := queue.Queue{}
	q.Push(2)
	q.Push(1)
	q.Push(3)
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
	fmt.Println(q.IsEmpty())
	q.Pop()
}
