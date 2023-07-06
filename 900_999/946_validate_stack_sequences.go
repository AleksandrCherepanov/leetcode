package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0, len(pushed))

	j := 0
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && j < len(popped) && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
			j++
		}
	}

	return len(stack) == 0
}
