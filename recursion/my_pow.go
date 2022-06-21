package recursion

func myPow(x float64, n int) float64 {
	negative := false
	if n < 0 {
		n *= -1
		negative = true
	}

	result := powMultiplication(x, n, x)

	if negative {
		result = 1 / result
	}

	return float64(int(result*100000)) / 100000
}

func powMultiplication(x float64, n int, prev float64) float64 {
	if n == 0 {
		return 1
	}

	return x * powMultiplication(x, n-1, prev)
}
