package main

func canVisitAllRooms(rooms [][]int) bool {
	seen := make(map[int]bool, 0)
	seen[0] = true

	stack := make([]int, 0, len(rooms))
	stack = append(stack, 0)

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for _, n := range rooms[node] {
			if _, ok := seen[n]; !ok {
				seen[n] = true
				stack = append(stack, n)
			}
		}
	}

	return len(seen) == len(rooms)
}
