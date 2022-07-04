package object

//结构体的定义
type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

// //结构体的方法:参数传递也是值传递,所以要通过方法改变对象内容的话要传指针，否则方法内对对象（拷贝）的改动是无效改动
// func (root *TreeNode) Traverse() {
// 	if root == nil {
// 		return
// 	}
// 	root.Left.Traverse()
// 	fmt.Println(root.Value)
// 	root.Right.Traverse()
// }

func (node *TreeNode) SetValue(value int) {
	node.Value = value
}

//工厂函数：实现C++中构造函数的效果
func CreateNode(value int) (node *TreeNode) {
	return &TreeNode{Value: value} //局部变量的地址也可以返回！如果不使用则会被垃圾回收，如果需要使用则程序不会认为此有异常
}

// func main() {
// //创建结构体对象
// // var root TreeNode
// // fmt.Println(root)

// root := TreeNode{value: 2}
// root.left = &TreeNode{} //注意root.left是指针，赋值时要取地址
// root.right = &TreeNode{value: 5}
// root.right.left = new(TreeNode)
// root.right.right = createNode(18)
// root.traverse()
// fmt.Println("------------")
// root.setValue(10)
// root.traverse()

// fmt.Println("------------")
// pRoot := &root
// pRoot.right.setValue(999)
// root.traverse()

// }
