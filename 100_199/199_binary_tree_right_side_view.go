package main

func rightSideView(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		count := len(q)
		for i := 0; i < count; i++ {
			node := q[0]
			q = q[1:]
			if i == count-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return result
}
