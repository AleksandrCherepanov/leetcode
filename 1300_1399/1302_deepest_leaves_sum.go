package main

func deepestLeavesSum(root *TreeNode) int {
	q := []*TreeNode{root}
	sum := 0

	for len(q) != 0 {
		l := len(q)
		sum = 0
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			sum += node.Val

			if node.Left != nil {
				q = append(q, node.Left)
			}

			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return sum
}
