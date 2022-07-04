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
