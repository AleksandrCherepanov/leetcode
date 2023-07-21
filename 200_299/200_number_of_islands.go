package main

import "fmt"

func numIslands(grid [][]byte) int {
	seen := make(map[int]bool)
	stack := make([][]int, 0)
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, ok := seen[i*400+j]; ok {
				continue
			}
			if grid[i][j] != '1' {
				continue
			}
			ans++

			stack = append(stack, []int{i, j})
			for len(stack) != 0 {
				el := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				ii := el[0]
				jj := el[1]

				if _, ok := seen[ii*400+jj]; ok {
					continue
				}

				seen[ii*400+jj] = true
				if ii-1 >= 0 && grid[ii-1][jj] == '1' {
					stack = append(stack, []int{ii - 1, jj})
				}

				if jj+1 < len(grid[ii]) && grid[ii][jj+1] == '1' {
					stack = append(stack, []int{ii, jj + 1})
				}

				if ii+1 < len(grid) && grid[ii+1][jj] == '1' {
					stack = append(stack, []int{ii + 1, jj})
				}

				if jj-1 >= 0 && grid[ii][jj-1] == '1' {
					stack = append(stack, []int{ii, jj - 1})
				}
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '1'},
		{'1', '1', '1', '0', '1'},
	}))
}
