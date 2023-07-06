package main

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, 0, len(asteroids))
	stack = append(stack, asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		cand := asteroids[i]
		for len(stack) != 0 && cand < 0 && stack[len(stack)-1] > 0 {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if cand*-1 < last {
				cand = last
				continue
			}

			if cand*-1 == last {
				cand = 0
			}
		}

		if cand != 0 {
			stack = append(stack, cand)
		}
	}

	return stack
}
