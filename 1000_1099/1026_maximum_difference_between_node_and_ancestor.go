package main

func maxAncestorDiff(root *TreeNode) int {
	return dfs(root, root.Val, root.Val)
}

func dfs(root *TreeNode, maxBefore int, minBefore int) int {
	if root == nil {
		return maxBefore - minBefore
	}

	newMax := maxBefore
	newMin := minBefore
	if root.Val > newMax {
		newMax = root.Val
	}
	if root.Val < newMin {
		newMin = root.Val
	}

	left := dfs(root.Left, newMax, newMin)
	right := dfs(root.Right, newMax, newMin)

	if left > right {
		return left
	}

	return right
}
