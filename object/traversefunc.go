package object

import "fmt"

func (node *TreeNode) Traverse() {
	node.TraverseFunc(func(tn *TreeNode) {
		fmt.Println(tn.Value)
	})
}
func (node *TreeNode) TraverseFunc(f func(*TreeNode)) {
	if node == nil {
		return
	}
	node.Left.TraverseFunc(f)
	f(node)
	node.Right.TraverseFunc(f)
}

func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(tn *TreeNode) {
			out <- tn
		})
		close(out) //func close(c chan<- Type)关闭通道，所有消息已发送完毕。
	}()
	return out
}
