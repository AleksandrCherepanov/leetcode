package main

func reachableNodes(n int, edges [][]int, restricted []int) int {
	graph := make(map[int][]int)
	seen := make(map[int]bool)

	for i := 0; i < len(restricted); i++ {
		seen[restricted[i]] = true
	}

	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]

		_, xok := seen[x]
		_, yok := seen[y]
		if xok || yok {
			continue
		}

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	stack := []int{0}
	ans := make([]int, 0, n)
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		seen[el] = true
		ans = append(ans, el)
		stack = stack[:len(stack)-1]

		for _, v := range graph[el] {
			if _, ok := seen[v]; !ok {
				stack = append(stack, v)
			}
		}
	}

	return len(ans)
}
