package main

func mostCompetitive(nums []int, k int) []int {
	stack := make([]int, 0, len(nums))

	for i, n := range nums {
		for len(stack) > 0 && k-len(stack) > len(nums)-i && stack[len(stack)-1] > n {
			stack = stack[:len(stack)-1]
		}

		if len(stack) < k {
			stack = append(stack, n)
		}
	}

	return stack
}
