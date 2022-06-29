package common

import "testing"

type mergeSortInput struct {
	name     string
	input    []int
	expected mergeSortResult
}

type mergeSortResult []int

var testCases = []mergeSortInput{
	{
		"case #1",
		[]int{1},
		[]int{1},
	},
	{
		"case #2",
		[]int{2, 7, 6, 3, 2, 3, 0, 9, 1},
		[]int{0, 1, 2, 2, 3, 3, 6, 7, 9},
	},
	{
		"case #3",
		[]int{1, 2, 3, 4, 5},
		[]int{1, 2, 3, 4, 5},
	},
	{
		"case #4",
		[]int{8, 7, 6, 5, 4, 3, 2, 1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8},
	},
}

func TestBottomUpIterative(t *testing.T) {
	for _, tc := range testCases {
		result := mergeSortResult(bottomUpIterative(tc.input))

		if len(result) != len(tc.expected) {
			mergeSortPrint(t, tc, result)
		}

		for i, n := range result {
			if n != tc.expected[i] {
				mergeSortPrint(t, tc, result)
			}
		}
	}
}

func TestTopDownRecursive(t *testing.T) {
	for _, tc := range testCases {
		result := mergeSortResult(topDownRecursive(tc.input))

		if len(result) != len(tc.expected) {
			mergeSortPrint(t, tc, result)
		}

		for i, n := range result {
			if n != tc.expected[i] {
				mergeSortPrint(t, tc, result)
			}
		}
	}
}

func mergeSortPrint(t *testing.T, tc mergeSortInput, result mergeSortResult) {
	t.Fatalf("%v. Expected: %v. Actual: %v", tc.name, tc.expected, result)
}
