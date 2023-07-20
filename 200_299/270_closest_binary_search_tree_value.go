package main

import "math"

func closestValue(root *TreeNode, target float64) int {
	q := []*TreeNode{root}
	min := math.Inf(0)
	result := math.Inf(0)

	for len(q) != 0 {
		n := q[0]
		q = q[1:]

		if n.Left != nil {
			q = append(q, n.Left)
		}

		if n.Right != nil {
			q = append(q, n.Right)
		}

		newMin := math.Abs(float64(n.Val) - target)
		if newMin == min {
			if result > float64(n.Val) {
				result = float64(n.Val)
			}
		}
		if newMin < min {
			min = newMin
			result = float64(n.Val)
		}
	}

	return int(result)
}
