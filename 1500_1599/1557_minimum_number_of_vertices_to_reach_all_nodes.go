package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	indegree := make(map[int]int, 0)

	for _, e := range edges {
		indegree[e[1]]++
	}

	result := make([]int, 0, len(indegree))
	for _, i := range indegree {
		if i == 0 {
			result = append(result, i)
		}
	}

	return result
}
