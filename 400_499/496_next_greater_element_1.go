package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	dict := make(map[int]int, len(nums2))
	answer := make([]int, len(nums1))

	for _, n := range nums2 {
		for len(stack) != 0 && stack[len(stack)-1] < n {
			dict[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, n)
	}

	for i, n := range nums1 {
		if v, ok := dict[n]; ok {
			answer[i] = v
		} else {
			answer[i] = -1
		}
	}

	return answer
}
