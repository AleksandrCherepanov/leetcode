package recursion

var cache map[int]int = make(map[int]int)

func fibonacci(n int) int {
	if n == 0 {
		cache[n] = 0
		return 0
	}

	if n == 1 {
		cache[n] = 1
		return 1
	}

	if _, ok := cache[n]; ok {
		return cache[n]
	}

	cache[n] = fibonacci(n-1) + fibonacci(n-2)
	return cache[n]
}
