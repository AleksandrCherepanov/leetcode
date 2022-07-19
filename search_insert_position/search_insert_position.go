package search_insert_position

func searchInsertPosition(nums []int, target int) int {
	low := 0
	high := len(nums)

	for low < high {
		mid := low + (high-low)/2
		pivot := nums[mid]

		if pivot == target {
			return mid
		}

		if pivot > target {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}
