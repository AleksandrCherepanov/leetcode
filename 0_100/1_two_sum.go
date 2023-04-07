package main

func twoSum(nums []int, target int) []int {
	traversed := map[int]int{nums[0]: 0}
	result := make([]int, 2)

	for i := 1; i < len(nums); i++ {
		diff := target - nums[i]
		if v, ok := traversed[diff]; ok {
			result[0] = v
			result[1] = i
			break
		}

		traversed[nums[i]] = i
	}

	return result
}
