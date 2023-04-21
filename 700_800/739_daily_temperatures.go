package main

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	answers := make([]int, len(temperatures))

	for i, t := range temperatures {
		for len(stack) != 0 && temperatures[stack[len(stack)-1]] < t {
			j := len(stack) - 1
			stack = stack[:len(stack)-1]
			answers[j] = i - j
		}

		stack = append(stack, i)
	}

	return answers
}
