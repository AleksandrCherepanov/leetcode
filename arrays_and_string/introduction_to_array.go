package arrays_and_string

import "fmt"

func pivotIndex(nums []int) int {
	sum := 0
	leftsum := 0

	//Time: O(n)
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	// Time: O(n)
	for i := 0; i < len(nums); i++ {
		if leftsum == sum-leftsum-nums[i] {
			return i
		} else {
			leftsum += nums[i]
		}
	}

	return -1
}

func dominantIndex(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	//Time O(n)
	maxIndex := 0
	max := nums[maxIndex]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
			maxIndex = i
		}
	}

	// Time O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] != max && max < nums[i]*2 {
			return -1
		}
	}

	return maxIndex
}

func plusOne(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	// Time O(n), Mem: O(n)
	reveresed := reverseSlice(nums)
	additional := 1

	// Time O(n)
	for i := 0; i < len(reveresed); i++ {
		reveresed[i] += additional
		additional = 0

		if reveresed[i] >= 10 {
			additional = reveresed[i] / 10
			reveresed[i] %= 10
		}
	}

	if additional != 0 {
		reveresed = append(reveresed, additional)
	}

	return reverseSlice(reveresed)
}

func reverseSlice(nums []int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp

		i++
		j--
	}

	return nums
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	verticalSteps := len(matrix) - 1
	horizontalSteps := len(matrix[0])

	result := make([]int, 0, len(matrix)*len(matrix[0]))

	i := 0
	j := 0

	h := true
	v := false
	hDirection := 1
	vDirection := 1

	for verticalSteps > 0 && horizontalSteps > 0 {
		if v && vDirection > 0 {
			for i < verticalSteps {
				result = append(result, matrix[i][j])
				i += vDirection
			}
			h = true
			v = false
			verticalSteps--
			i = verticalSteps
			vDirection = -1
		}

		if v && vDirection < 0 {
			for i > 0 {
				result = append(result, matrix[i][j])
				i += vDirection
			}

			h = true
			v = false
			verticalSteps--
			i = len(matrix) - verticalSteps
			vDirection = 1
		}

		if h && hDirection > 0 {
			for j < horizontalSteps {
				fmt.Println(j)
				result = append(result, matrix[i][j])
				j += hDirection
			}

			h = false
			v = true
			horizontalSteps--
			j = horizontalSteps - 1
			hDirection = -1
		}

		if h && hDirection < 0 {
			for j > 0 {
				fmt.Println(j)
				result = append(result, matrix[i][j])
				j += hDirection
			}

			h = false
			v = true
			horizontalSteps--
			j = len(matrix[0]) - horizontalSteps
			hDirection = 1
		}
	}

	return result
}
