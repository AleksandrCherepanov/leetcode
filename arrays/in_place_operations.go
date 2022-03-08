package arrays

func isEven(num int) bool {
	return num%2 == 0
}

func SortArrayByParity(nums []int) []int {
	i := 0
	j := 0

	for j < len(nums) && i < len(nums) {
		if !isEven(nums[j]) {
			for i < len(nums) && !isEven(nums[i]) {
				i++
			}
			if i == len(nums) {
				break
			}
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			j++
			continue
		}
		j++
		i++
	}

	return nums
}

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	count := 0
	for i <= j {
		if nums[j] == val {
			nums[j] = -1
			j--
			count++
			continue
		}

		if nums[i] == val {
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
			nums[j] = -1
			j--
			i++
			count++
			continue
		}

		i++
	}

	return len(nums) - count
}
