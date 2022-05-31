package arrays_and_string

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	m := len(matrix)
	n := len(matrix[0])

	result := make([]int, 0, m*n)
	step := 0
	i := 0
	j := 0
	si := m - 1
	sj := n
	mi := 1
	mj := 1
	c := 0

	for step < m*n {
		c = 0
		for ; c < sj; c++ {
			result = append(result, matrix[i][j])
			j += mj
			step++
		}
		j -= mj
		i += mi

		c = 0
		for ; c < si; c++ {
			result = append(result, matrix[i][j])
			i += mi
			step++
		}

		mi *= -1
		mj *= -1
		i += mi
		j += mj
		si--
		sj--
	}

	return result
}
