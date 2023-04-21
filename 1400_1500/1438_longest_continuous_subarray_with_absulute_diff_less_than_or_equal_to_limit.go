package main

import (
	"math"
)

func longestSubarray(nums []int, limit int) int {
	minQ := make([]int, 0, len(nums))
	maxQ := make([]int, 0, len(nums))

	left := 0
	ans := 0

	for right, n := range nums {
		for len(minQ) > 0 && n < minQ[len(minQ)-1] {
			minQ = minQ[:len(minQ)-1]
		}
		for len(maxQ) > 0 && n > maxQ[len(maxQ)-1] {
			maxQ = maxQ[:len(maxQ)-1]
		}

		minQ = append(minQ, n)
		maxQ = append(maxQ, n)

		for math.Abs(float64(minQ[0]-maxQ[0])) > float64(limit) {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}

	return ans
}
