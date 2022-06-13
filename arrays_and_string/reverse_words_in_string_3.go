package arrays_and_string

func reverseWords3(s string) string {
	space := byte(' ')
	j := 0
	result := []byte(s)

	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			k := (i - j + 1) / 2
			for c := 0; c < k; c++ {
				t := result[j+c]
				result[j+c] = result[i-c]
				result[i-c] = t
			}
		}
		if result[i] == space {
			k := (i - j) / 2
			for c := 0; c < k; c++ {
				t := result[j+c]
				result[j+c] = result[i-c-1]
				result[i-c-1] = t
			}
			j = i + 1
		}
	}

	return string(result)
}
