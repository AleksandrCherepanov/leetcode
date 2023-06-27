package main

// T O(n) S O(n^2)
func canSeePersonsCount(heights []int) []int {
	stack := make([]int, 0, len(heights))
	list := make([][]int, len(heights))
	answer := make([]int, 0, len(heights))

	for i, n := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] < n {
			lastElement := stack[len(stack)-1]
			if len(list[lastElement]) == 0 || list[lastElement][len(list[lastElement])-1] < n {
				list[lastElement] = append(list[lastElement], n)
			}
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			list[stack[len(stack)-1]] = append(list[stack[len(stack)-1]], n)
		}
		stack = append(stack, i)
	}

	for _, el := range list {
		answer = append(answer, len(el))
	}

	return answer
}

// T O(n) S O(n)
func canSeePersonsCountOpt(heights []int) []int {
	stack := make([]int, 0, len(heights))
	answer := make([]int, len(heights))

	for i, n := range heights {
		for len(stack) != 0 && heights[stack[len(stack)-1]] < n {
			answer[stack[len(stack)-1]]++
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			answer[stack[len(stack)-1]]++
		}

		stack = append(stack, i)
	}

	return answer
}
