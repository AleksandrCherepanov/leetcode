package arrays_and_string

func reverseString(s []byte) {
	i := 0
	j := len(s) - 1

	for i < j {
		t := s[i]
		s[i] = s[j]
		s[j] = t

		i++
		j--
	}
}
