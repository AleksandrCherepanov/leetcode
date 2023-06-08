package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if targetSum == 0 || root == nil {
		return false
	}

	stack := make([]Pair, 0)

	stack = append(stack, Pair{Node: root, Subsum: 0})

	for len(stack) != 0 {
		pair := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if pair.Node.Left == nil && pair.Node.Right == nil {
			return pair.Subsum+pair.Node.Val == targetSum
		}

		pair.Subsum += pair.Node.Val
		if pair.Node.Left != nil {
			stack = append(stack, Pair{Node: pair.Node.Left, Subsum: pair.Subsum})
		}
		if pair.Node.Right != nil {
			stack = append(stack, Pair{Node: pair.Node.Right, Subsum: pair.Subsum})
		}
	}

	return false
}

func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
	return dfs(root, 0, targetSum)
}

func dfs(root *TreeNode, curr int, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return curr+root.Val == targetSum
	}

	curr += root.Val
	left := dfs(root.Left, curr, targetSum)
	right := dfs(root.Right, curr, targetSum)

	return left || right
}
