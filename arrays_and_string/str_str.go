package arrays_and_string

func strStr(haystack, needle string) int {
	if haystack == "" {
		return -1
	}

	if needle == "" {
		return 0
	}

	index := -1
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		found := true
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				found = false
				break
			}
		}

		if found {
			index = i
			break
		}
	}

	return index
}
