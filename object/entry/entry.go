package main

import (
	"fmt"

	"coolgiserz.com/learngo/object"
)

//扩展已有类型：通过组合的方式扩展
type myTreeNode struct {
	myNode *object.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.myNode == nil {
		return
	}
	left := myTreeNode{myNode.myNode.Left}
	right := myTreeNode{myNode.myNode.Right}
	left.postOrder()
	right.postOrder()
	fmt.Println(myNode.myNode.Value)
}
func main() {
	//创建结构体对象
	// var root TreeNode
	// fmt.Println(root)

	root := object.TreeNode{Value: 2}
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

	fmt.Println("---test traversefunc---------")
	nodeCount := 0
	root.TraverseFunc(func(tn *object.TreeNode) { //统计节点数
		nodeCount += 1
	})
	fmt.Println("Nodecount: ", nodeCount)
	fmt.Println("-----+++-------")
	myroot := myTreeNode{&root}
	myroot.postOrder()
}
