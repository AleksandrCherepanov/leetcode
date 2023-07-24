package main

import "strings"

func makeGood(s string) string {
	runes := []rune(s)

	stack := []string{}
	for i := 0; i < len(runes); i++ {
		lenStack := len(stack)
		curr := string(runes[i])
		if lenStack == 0 {
			stack = append(stack, curr)
			continue
		}

		prev := stack[lenStack-1]
		if prev != curr && strings.ToLower(prev) == strings.ToLower(curr) {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, curr)
		}
	}

	return strings.Join(stack, "")
}
