package recursion

import "math"

func kthGrammar(n int, k int) int {
	if n == 1 && k == 1 {
		return 0
	}

	mid := int(math.Pow(2, float64(n-1)) / 2)

	if k <= mid {
		return kthGrammar(n-1, k)
	} else {
		res := kthGrammar(n-1, k-mid)
		if res == 0 {
			return 1
		}

		return 0
	}
}
