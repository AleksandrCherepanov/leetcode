package arrays_and_string

import "fmt"

func pascalTriangle2(rowIndex int) []int {
	result := [][]int{{1}}

	if rowIndex >= 1 {
		result = append(result, []int{1, 1})
	}

	for i := 2; i <= rowIndex; i++ {
		prevLine := i - 1
		line := make([]int, 0, len(result[prevLine])+1)
		for j := 0; j < i; j++ {
			if j-1 < 0 {
				line = append(line, result[prevLine][j])
				continue
			} else {
				line = append(line, result[prevLine][j]+result[prevLine][j-1])
			}

			if j+1 >= i {
				line = append(line, result[prevLine][len(result[prevLine])-1])
				continue
			}
		}
		result = append(result, line)
	}

	fmt.Println(result)

	return result[len(result)-1]
}
