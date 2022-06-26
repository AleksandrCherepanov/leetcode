package recursion

func uniqueBstGenerate(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}

	return treeGenereate(1, n)
}

func treeGenereate(start int, end int) []*TreeNode {
	nodes := []*TreeNode{}
	if start > end {
		nodes = append(nodes, nil)
		return nodes
	}

	for i := start; i <= end; i++ {
		left := treeGenereate(start, i-1)
		right := treeGenereate(i+1, end)

		for _, l := range left {
			for _, r := range right {
				nodes = append(nodes, createTreeNode(i, l, r))
			}
		}
	}

	return nodes
}
