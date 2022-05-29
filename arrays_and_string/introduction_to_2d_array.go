package arrays_and_string

func DiagonalTraverse(mat [][]int) []int {
	if len(mat) == 0 {
		return []int{}
	}

	count := len(mat) * len(mat[0])

	result := make([]int, 0, count)
	i := 0
	j := 0
	up := 1

	for count > 0 {
		if i >= len(mat) {
			i--
			j += 2
			up *= -1
		}

		if j >= len(mat[0]) {
			j--
			i += 2
			up *= -1
		}

		if i < 0 {
			i++
			up *= -1
		}

		if j < 0 {
			j++
			up *= -1
		}

		result = append(result, mat[i][j])

		if up > 0 {
			i--
			j++
		} else {
			i++
			j--
		}

		count--
	}

	return result
}
