package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return minDepth(root.Right) + 1
	}

	if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left < right {
		return left + 1
	}

	return right + 1
}
