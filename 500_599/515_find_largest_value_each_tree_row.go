package main

func largestValues(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) != 0 {
		count := len(q)
		maxNode := q[0].Val
		for i := 0; i < count; i++ {
			node := q[0]
			q = q[1:]

			if node.Val > maxNode {
				maxNode = node.Val
			}

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		res = append(res, maxNode)
	}

	return res
}
