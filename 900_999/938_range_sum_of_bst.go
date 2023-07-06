package main

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	sum := 0
	if low <= root.Val && root.Val <= high {
		sum += root.Val
	}

	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}

	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
