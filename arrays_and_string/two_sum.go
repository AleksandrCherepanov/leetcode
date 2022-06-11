package arrays_and_string

func twoSum(numbers []int, target int) []int {
	result := make([]int, 2)

	i := 0
	j := len(numbers) - 1

	for i < j {
		sum := numbers[i] + numbers[j]
		if sum > target {
			j--
		}
		if sum < target {
			i++
		}
		if sum == target {
			result[0] = i + 1
			result[1] = j + 1
			break
		}
	}

	return result
}
