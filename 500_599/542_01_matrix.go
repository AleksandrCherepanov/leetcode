package main

func updateMatrix(mat [][]int) [][]int {
	m := len(mat)
	n := len(mat[0])

	seen := make([][]bool, 0, m)
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	queue := make([][]int, 0, m)

	for i := 0; i < m; i++ {
		seenRow := make([]bool, 0, n)
		for j := 0; j < n; j++ {
			seenRow = append(seenRow, mat[i][j] == 0)
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j, 0})
			}
		}
		seen = append(seen, seenRow)
	}

	for len(queue) > 0 {
		l := len(queue)
		for q := 0; q < l; q++ {
			el := queue[0]
			queue = queue[1:]

			for _, d := range directions {
				newX := el[0] + d[0]
				newY := el[1] + d[1]

				if valid(newX, newY, m, n) && !seen[newX][newY] {
					distance := el[2] + 1
					queue = append(queue, []int{newX, newY, distance})
					seen[newX][newY] = true
					mat[newX][newY] = distance
				}
			}
		}
	}

	return mat
}

func valid(i, j, m, n int) bool {
	return 0 <= i && 0 <= j && i < m && j < n
}
