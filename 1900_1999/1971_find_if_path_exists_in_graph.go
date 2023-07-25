package main

func validPath(n int, edges [][]int, source int, destination int) bool {
	if source == destination {
		return true
	}

	graph := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	seen := make(map[int]bool)
	seen[source] = true

	stack := make([]int, 0, len(graph))
	stack = append(stack, source)

	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, v := range graph[n] {
			if v == destination {
				return true
			}
			if _, ok := seen[v]; !ok {
				seen[v] = true
				stack = append(stack, v)
			}
		}
	}

	return false
}
