package main

import (
	"fmt"

	"coolgiserz.com/learngo/object"
)

//通过内嵌扩展已有类型
type myTreeNode struct {
	*object.TreeNode //内嵌
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.TreeNode == nil {
		return
	}
	left := myTreeNode{myNode.Left}
	right := myTreeNode{myNode.Right}
	left.postOrder()
	right.postOrder()
	fmt.Println(myNode.Value)
}
func main() {
	//创建结构体对象
	// var root TreeNode
	// fmt.Println(root)

	root := myTreeNode{&object.TreeNode{Value: 2}}
	root.Left = &object.TreeNode{} //注意root.left是指针，赋值时要取地址
	root.Right = &object.TreeNode{Value: 5}
	root.Right.Left = new(object.TreeNode)
	root.Right.Right = object.CreateNode(18)
	root.Traverse()
	fmt.Println("------------")
	root.SetValue(10)
	root.Traverse()

	// fmt.Println("------------")
	// pRoot := &root
	// pRoot.Right.SetValue(999)
	// root.Traverse()

	fmt.Println("-----嵌入类型+++-------")
	myroot := root
	myroot.postOrder()
}
