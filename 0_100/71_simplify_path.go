package main

import "strings"

func simplifyPath(path string) string {
	tokens := strings.Split(path, "/")
	stack := []string{}

	for _, token := range tokens {
		if token == "." || token == "/" || token == "" {
			continue
		}

		if token == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, token)
		}
	}

	return "/" + strings.Join(stack, "/")
}
