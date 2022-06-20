package recursion

var climbingStairsCache map[int]int = make(map[int]int)

func climbingStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		climbingStairsCache[n] = n
		return n
	}

	if _, ok := climbingStairsCache[n]; ok {
		return climbingStairsCache[n]
	}

	climbingStairsCache[n] = climbingStairs(n-1) + climbingStairs(n-2)
	return climbingStairsCache[n]
}
