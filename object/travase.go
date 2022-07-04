package object

// import "fmt"

// //为结构体定义的方法可以写在不同的文件中
// //结构体的方法:参数传递也是值传递,所以要通过方法改变对象内容的话要传指针，否则方法内对对象（拷贝）的改动是无效改动
// func (root *TreeNode) Traverse() {
// 	if root == nil {
// 		return
// 	}
// 	root.Left.Traverse()
// 	fmt.Println(root.Value)
// 	root.Right.Traverse()
// }
