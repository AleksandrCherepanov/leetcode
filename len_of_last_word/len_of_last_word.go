package len_od_last_word

func lenOfLastWord(s string) int {
	lenght := 0
	for i := len(s) - 1; i >= 0; i-- {
		if lenght != 0 && s[i] == byte(' ') {
			return lenght
		}

		if s[i] != byte(' ') {
			lenght++
		}
	}

	return lenght
}
