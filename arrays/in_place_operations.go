package arrays

import (
	"sort"
)

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

func heightCheck(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)

	result := 0
	for i := 0; i < len(tmp); i++ {
		if tmp[i] != nums[i] {
			result++
		}
	}

	return result
}

func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	sort.Ints(nums)
	max := nums[len(nums)-1]
	count := 0

	for i := len(nums) - 2; i >= 0; i-- {
		if count == 2 {
			break
		}

		if nums[i] != max {
			max = nums[i]
			count++
		}
	}

	if count < 2 {
		return nums[len(nums)-1]
	} else {
		return max
	}
}
