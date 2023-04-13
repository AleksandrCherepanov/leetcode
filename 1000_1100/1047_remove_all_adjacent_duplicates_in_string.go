package main

import "strings"

func removeDuplicates(s string) string {
	stack := []string{}

	for _, c := range (s) {
		cs := string(c)
		if len(stack) != 0 {
			if stack[len(stack) - 1] == cs {
				stack = stack[:len(stack) - 1]
				continue
			} 
		} 

		stack = append(stack, cs)
	}

	return strings.Join(stack, "")
}
