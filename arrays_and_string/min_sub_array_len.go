package arrays_and_string

func minSubArrayLen(target int, nums []int) int {
	sum := 0
	result := 0
	j := 0
	i := 0
	found := false
	for i < len(nums) {
		sum += nums[i]
		for sum >= target && j <= i {
			if result == 0 || i-j+1 < result {
				result = i - j + 1
			}
			if !found {
				found = true
			}
			sum -= nums[j]
			j++
		}

		i++
	}

	if !found {
		return 0
	}

	return result
}
