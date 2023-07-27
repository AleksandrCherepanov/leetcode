package main

import "fmt"

func maxAreaOfIsland(grid [][]int) int {
	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	seen := make([][]bool, 0)
	for x := 0; x < len(grid); x++ {
		row := make([]bool, 0, len(grid[x]))
		for y := 0; y < len(grid[x]); y++ {
			row = append(row, false)
		}
		seen = append(seen, row)
	}
	stack := make([][]int, 0)
	ans := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			if grid[x][y] == 0 {
				continue
			}

			if seen[x][y] {
				continue
			}

			stack = append(stack, []int{x, y})
			seen[x][y] = true
			counter := 1
			for len(stack) > 0 {
				el := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				for _, d := range directions {
					newX := el[0] + d[0]
					newY := el[1] + d[1]

					if isCorrectPoint(newX, newY, len(grid), len(grid[el[0]])) && grid[newX][newY] == 1 && !seen[newX][newY] {
						stack = append(stack, []int{newX, newY})
						seen[newX][newY] = true
						counter++
					}
				}
			}

			if counter > ans {
				ans = counter
			}
		}
	}

	return ans
}

func isCorrectPoint(x, y, h, w int) bool {
	return x >= 0 && x < h && y >= 0 && y < w
}

func main() {
	grid := [][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}
	fmt.Println(maxAreaOfIsland(grid))
}
