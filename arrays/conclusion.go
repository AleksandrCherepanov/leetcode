package arrays

import "sort"

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

func findDisappearedNumbers(nums []int) []int {
	n := make([]int, len(nums))
	disappeared := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		n[nums[i]-1] = nums[i]
	}

	for i, v := range n {
		if v == 0 {
			disappeared = append(disappeared, i+1)
		}
	}

	return disappeared
}

func sortedSquares(nums []int) []int {
	negativeSquares := make([]int, 0, len(nums)) //Mem: O(n)
	positiveSquares := make([]int, 0, len(nums)) //Mem: O(n)
	result := make([]int, 0, len(nums))          //Mem: O(n)

	// Time: O(n)
	for _, n := range nums {
		if n < 0 {
			negativeSquares = append(negativeSquares, n*n)
		} else {
			positiveSquares = append(positiveSquares, n*n)
		}
	}

	i := 0
	j := len(negativeSquares) - 1

	//Time: O(n)
	for i < len(positiveSquares) && j >= 0 {
		if positiveSquares[i] < negativeSquares[j] {
			result = append(result, positiveSquares[i])
			i++
		} else {
			result = append(result, negativeSquares[j])
			j--
		}
	}

	if j < 0 && i < len(positiveSquares) {
		result = append(result, positiveSquares[i:]...)
	}

	if i == len(positiveSquares) && j >= 0 {
		//Time: O(n)
		for j >= 0 {
			result = append(result, negativeSquares[j])
			j--
		}
	}

	return result
}
