package main

func lengthOfLongestSubstring(s string) int {
	dict := map[byte]int{s[0]: 1}

	maxLen := 1
	currLen := 0
	l := 0
	r := 1
	for r < len(s) {
		dict[s[r]]++
		for dict[s[r]] > 1 {
			dict[s[l]]--
			l++
		}
		currLen = r - l + 1
		if currLen > maxLen {
			maxLen = currLen
		}

		r++
	}

	return maxLen
}
