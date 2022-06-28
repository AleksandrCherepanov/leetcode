package common

func bottomUpIterative(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	result := make([][]int, 0, len(nums))
	for _, n := range nums {
		result = append(result, []int{n})
	}

	i := 0
	j := 0
	for len(result) != 1 {

	}
}

func topDownRecursive(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
