package binary_tree_inorder_traversal

import "github.com/AleksandrCherepanov/leetcode/structure"

func inorderTraversal(root *structure.TreeNode) []int {
	result := make([]int, 0, 0)

	traverse(root, &result)

	return result
}

func traverse(root *structure.TreeNode, slice *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, slice)
	*slice = append(*slice, root.Val)
	traverse(root.Right, slice)
}
