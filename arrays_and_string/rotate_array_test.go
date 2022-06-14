package arrays_and_string

import "testing"

type tRotateArrayResult []int

type tRotateArrayCase struct {
	name     string
	a        []int
	b        int
	expected tRotateArrayResult
}

func TestRotateArray(t *testing.T) {
	testCases := []tRotateArrayCase{
		{
			"Case #1",
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			"Case #2",
			[]int{-1, -100, 3, 99},
			2,
			[]int{3, 99, -1, -100},
		},
		{
			"Case #3",
			[]int{1, 2, 3},
			4,
			[]int{3, 1, 2},
		},
	}

	for _, testCase := range testCases {
		rotateArray(testCase.a, testCase.b)
		if len(testCase.a) != len(testCase.expected) {
			rotateArrayPrintError(t, testCase, testCase.a)
		}

		for i, n := range testCase.expected {
			if n != testCase.a[i] {
				rotateArrayPrintError(t, testCase, testCase.a)
			}
		}
	}
}

func rotateArrayPrintError(t *testing.T, testCase tRotateArrayCase, result tRotateArrayResult) {
	t.Fatalf(
		"Case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
