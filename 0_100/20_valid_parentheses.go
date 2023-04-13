package main

func isValid(s string) bool {
	parentheses := map[string]string{"(": ")", "{": "}", "[": "]"}

	stack := []string{}
	for _, c := range(s) {
		cs := string(c)
		if _, ok := parentheses[cs]; ok {
			stack = append(stack, cs)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack) - 1]
			if cs != parentheses[top] {
				return false
			}
			stack = stack[:len(stack) - 1]
		}
	}

	return len(stack) == 0
}
