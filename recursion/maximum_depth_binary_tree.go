package recursion

var depth int

func maximumDepthBinaryTree(root *TreeNode) int {
	depth = 0

	if root == nil {
		return depth
	}

	treeWalker(root, depth)

	return depth
}

func treeWalker(root *TreeNode, counter int) {
	if root == nil {
		if counter > depth {
			depth = counter
		}
		return
	}

	treeWalker(root.Left, counter+1)
	treeWalker(root.Right, counter+1)
}
