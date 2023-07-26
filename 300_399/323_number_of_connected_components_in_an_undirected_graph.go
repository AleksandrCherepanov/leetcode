package main

func countComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)

	for i := 0; i < len(edges); i++ {
		x := edges[i][0]
		y := edges[i][1]

		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	ans := 0
	seen := make(map[int]bool)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if _, ok := seen[i]; !ok {
			ans++
			stack = append(stack, i)
		}

		for len(stack) > 0 {
			el := stack[len(stack)-1]
			seen[el] = true
			stack = stack[:len(stack)-1]
			for _, v := range graph[el] {
				if _, ok := seen[v]; !ok {
					stack = append(stack, v)
				}
			}
		}
	}

	return ans
}
