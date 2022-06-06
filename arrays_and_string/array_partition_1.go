package arrays_and_string

import "sort"

func arrayPartition1(nums []int) int {
	result := 0

	sort.Ints(nums)

	i := len(nums) - 1
	j := len(nums) - 2

	for j >= 0 {
		if nums[i] < nums[j] {
			result += nums[i]
		} else {
			result += nums[j]
		}
		i -= 2
		j -= 2
	}

	return result
}
