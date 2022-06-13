package arrays_and_string

func rotateArray(nums []int, k int) {
	if len(nums) == 1 || len(nums) == k {
		return
	}

	k = k % len(nums)
	index := len(nums) - k

	revert(nums, 0, index-1)
	revert(nums, index, len(nums)-1)
	revert(nums, 0, len(nums)-1)
}

func revert(a []int, start, end int) {
	for start < end {
		t := a[start]
		a[start] = a[end]
		a[end] = t
		start++
		end--
	}
}
