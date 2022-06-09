package arrays_and_string

func longestCommonPrefix(strs []string) string {
	result := ""

	if len(strs) == 0 {
		return result
	}

	i := 0
	search := true

	for search {
		if i == len(strs[0]) {
			break
		}

		letter := strs[0][i]
		for _, word := range strs {
			if i == len(word) {
				search = false
				break
			}

			if word[i] != letter {
				search = false
				break
			}
		}

		if search {
			result += string(letter)
		}
		i++
	}

	return result
}
