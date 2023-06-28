package main

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter = 0
	maxDepth(root, &diameter)
	return diameter
}

func maxDepth(root *TreeNode, diameter *int) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left, diameter)
	right := maxDepth(root.Right, diameter)

	if left+right > *diameter {
		*diameter = left + right
	}

	if left > right {
		return left + 1
	}

	return right + 1
}
