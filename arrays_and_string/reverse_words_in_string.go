package arrays_and_string

import (
	"strings"
)

func reverseWordsInString(s string) string {
	newString := make([]byte, 0, len([]byte(s)))

	space := 0
	for i := 0; i < len([]byte(s)); i++ {
		if s[i] == byte(' ') {
			space++
		} else {
			space = 0
		}

		if space > 1 {
			continue
		}

		newString = append(newString, s[i])
	}

	s = strings.TrimSpace(string(newString))

	words := strings.Split(s, " ")
	i := 0
	j := len(words) - 1

	for i < j {
		t := words[i]
		words[i] = words[j]
		words[j] = t
		i++
		j--
	}

	return strings.Join(words, " ")
}
