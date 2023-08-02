package main

import "strings"

func removeStars(s string) string {
	stack := make([]string, 0, len(s))

	for _, r := range s {
		stack = append(stack, string(r))
		for len(stack) != 0 && stack[len(stack)-1] == "*" {
			stack = stack[:len(stack)-2]
		}
	}

	return strings.Join(stack, "")
}
