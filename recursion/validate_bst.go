package recursion

func validateBst(root *TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(root *TreeNode, low *TreeNode, high *TreeNode) bool {
	if root == nil {
		return true
	}

	if (low != nil && root.Val <= low.Val) || (high != nil && root.Val >= high.Val) {
		return false
	}

	left := validate(root.Left, low, root)
	right := validate(root.Right, root, high)

	return left && right
}
