package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	insert(root, val)
	return root
}

func insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if val > root.Val {
		root.Right = insert(root.Right, val)
	} else {
		root.Left = insert(root.Left, val)
	}

	return root
}
