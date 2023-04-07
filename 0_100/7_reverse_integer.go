package main

import "math"

func reverse(x int) int {
	result := 0
	negative := false
	if x < 0 {
		negative = true
		x *= -1
	}

	for x != 0 {
		if result * 10 > math.MaxInt32 {
			return 0
		}

		result *= 10
		n := x % 10
		if result + n > math.MaxInt32 {
			return 0
		}

		result += n
		x /= 10
	}

	if negative {
		if result * -1 > math.MaxInt32 {
			return 0
		}
		result *= -1
	}

	return result
}
