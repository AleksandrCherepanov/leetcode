package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	q := []*TreeNode{root}
	lr := true
	for len(q) != 0 {
		l := len(q)
		lvl := make([]int, l)
		for i := 0; i < l; i++ {
			n := q[0]
			q = q[1:]

			if lr {
				lvl[i] = n.Val
			} else {
				lvl[l-i-1] = n.Val
			}

			if n.Left != nil {
				q = append(q, n.Left)
			}

			if n.Right != nil {
				q = append(q, n.Right)
			}
		}

		lr = !lr
		ans = append(ans, lvl)
	}

	return ans
}
