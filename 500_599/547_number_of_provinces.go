package main

func gdfs(node int, graph map[int][]int) {
	for _, n := range graph[node] {
		if _, ok := seen[n]; !ok {
			seen[n] = true
			gdfs(n, graph)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	graph := make(map[int][]int)
	seen = make(map[int]bool)
	ans := 0

	for i := 0; i < len(isConnected); i++ {
		for j := i + 1; j < len(isConnected[i]); j++ {
			if isConnected[i][j] == 1 {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	for i := 0; i < len(isConnected); i++ {
		if _, ok := seen[i]; !ok {
			ans++
			seen[i] = true
			gdfs(i, graph)
		}
	}

	return ans
}
