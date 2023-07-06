package main

type Ans struct {
	ans []int
}

func (a *Ans) Add(v int) {
	a.ans = append(a.ans, v)
}

func dfs(root *TreeNode, ans *Ans) {
	if root == nil {
		return
	}

	dfs(root.Left, ans)
	ans.Add(root.Val)
	dfs(root.Right, ans)
}

func getMinimumDifference(root *TreeNode) int {
	ans := &Ans{ans: make([]int, 0)}
	dfs(root, ans)

	min := 100000
	for i := 0; i < len(ans.ans)-1; i++ {
		newVal := ans.ans[i] - ans.ans[i+1]
		if newVal < 0 {
			newVal *= -1
		}

		if newVal < min {
			min = newVal
		}
	}

	return min
}
