package main

import (
	"strconv"
)

func minReorder(n int, connections [][]int) int {
	graph := make(map[int][]int, 0)
	roads := make(map[string]bool, 0)
	for i := 0; i < len(connections); i++ {
		x := connections[i][0]
		y := connections[i][1]

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
		roads[getHash(x, y)] = true
	}

	seen := make(map[int]bool, 0)
	seen[0] = true

	stack := make([]int, 0, n)
	stack = append(stack, 0)
	ans := 0

	for len(stack) != 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, v := range graph[el] {
			if _, ok := seen[v]; !ok {
				if _, ok := roads[getHash(el, v)]; ok {
					ans++
				}
				seen[v] = true
				stack = append(stack, v)
			}
		}
	}

	return ans
}

func getHash(x, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}
