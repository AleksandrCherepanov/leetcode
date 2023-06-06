package main

func getAverages(nums []int, k int) []int {
	ans := make([]int, len(nums))
	l := 0
	r := 0

	sum := 0
	count := 0
	for r < len(nums) && l < len(nums) {
		ans[r] = -1
		sum += nums[r]
		count++

		for r-l == k*2 {
			ans[r-k] = sum / count
			sum -= nums[l]
			count--
			l++
		}

		r++
	}

	return ans
}
