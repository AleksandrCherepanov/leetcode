package structure

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  left,
		Right: right,
	}
}

func (root *TreeNode) PrintPreOrder() {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	root.Left.PrintPreOrder()
	root.Right.PrintPreOrder()
}

func (root *TreeNode) PrintInOrder() {
	if root == nil {
		return
	}

	root.Left.PrintInOrder()
	fmt.Println(root.Val)
	root.Right.PrintInOrder()
}
