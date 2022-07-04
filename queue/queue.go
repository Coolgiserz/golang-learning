package queue

//通过别名扩展已有类型
type Queue []int

func (q *Queue) Push(value int) {
	*q = append(*q, value)
}

func (q *Queue) Pop() int {
	if len(*q) == 0 {
		panic("length of queue is zero!")
	}
	head := (*q)[0]
	(*q) = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
