package main

func backspaceCompare(s string, t string) bool {
	first := processBackspace(s)
	second := processBackspace(t)

	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func processBackspace(s string) []rune {
	result := []rune{}

	for _, c := range (s) {
		if c != rune('#') {
			result = append(result, c)	
		} else if len(result) != 0 {
			result = result[:len(result)-1]
		} 

	}

	return result
}
