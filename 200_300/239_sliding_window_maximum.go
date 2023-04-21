package main

func maxSlidingWindow(nums []int, k int) []int {
	q := make([]int, 0)
	ans := make([]int, 0, len(nums))

	for i, n := range nums {
		for len(q) > 0 && n > nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}

		q = append(q, i)

		if q[0]+k == i {
			q = q[1:]
		}

		if i >= k-1 {
			ans = append(ans, nums[q[0]])
		}
	}

	return ans
}
