package queue

import "fmt"

func ExampleQueue_Pop() {
	q := &Queue{1}
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	//Output:
	//1
	//2
}
