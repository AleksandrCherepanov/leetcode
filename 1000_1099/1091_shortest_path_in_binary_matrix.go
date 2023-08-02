package main

import "strconv"

func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] == 1 {
		return -1
	}

	n := len(grid)
	queue := make([][]int, 0, len(grid))
	seen := make(map[string]bool)
	directions := [][]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}
	ans := 0

	queue = append(queue, []int{0, 0})
	seen[hash(0, 0)] = true
	for len(queue) > 0 {
		l := len(queue)
		ans++

		for i := 0; i < l; i++ {
			el := queue[0]
			queue = queue[1:]
			if el[0] == n-1 && el[1] == n-1 {
				return ans
			}

			for _, d := range directions {
				newX := el[0] + d[0]
				newY := el[1] + d[1]

				valid := check(newX, newY, n, grid)
				_, hasSeen := seen[hash(newX, newY)]
				if valid && !hasSeen {
					queue = append(queue, []int{newX, newY})
					seen[hash(newX, newY)] = true
				}
			}
		}
	}

	return -1
}

func hash(i, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func check(i, j, n int, grid [][]int) bool {
	return i >= 0 && i < n && j >= 0 && j < n && grid[i][j] == 0
}
