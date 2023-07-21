package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return search(root, -100000)
}

func search(root *TreeNode, maxVal int) int {
	if root == nil {
		return 0
	}

	if root.Val > maxVal {
		maxVal = root.Val
	}

	left := search(root.Left, maxVal)
	right := search(root.Right, maxVal)
	ans := left + right
	if root.Val >= maxVal {
		ans = ans + 1
	}

	return ans
}
