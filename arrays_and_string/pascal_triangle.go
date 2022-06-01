package arrays_and_string

func Generate(numRows int) [][]int {
	result := [][]int{{1}}

	if numRows > 1 {
		result = append(result, []int{1, 1})
	}

	for i := 3; i <= numRows; i++ {
		prevLine := i - 2
		line := make([]int, 0, len(result[prevLine])+1)
		for j := 0; j < i-1; j++ {
			if j-1 < 0 {
				line = append(line, result[prevLine][j])
				continue
			} else {
				line = append(line, result[prevLine][j]+result[prevLine][j-1])
			}

			if j+1 >= i-1 {
				line = append(line, result[prevLine][len(result[prevLine])-1])
				continue
			}
		}
		result = append(result, line)
	}

	return result
}
