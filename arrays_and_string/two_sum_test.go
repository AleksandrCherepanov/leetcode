package arrays_and_string

import "testing"

type tTwoSumOrderResult []int

type tTwoSumOrderCase struct {
	name     string
	a        []int
	b        int
	expected tTwoSumOrderResult
}

func TestTwoSumOrder(t *testing.T) {
	testCases := []tTwoSumOrderCase{
		{
			"Case #1",
			[]int{2, 7, 11, 15},
			9,
			[]int{1, 2},
		},
		{
			"Case #2",
			[]int{2, 3, 4},
			6,
			[]int{1, 3},
		},
		{
			"Case #3",
			[]int{-1, 0},
			-1,
			[]int{1, 2},
		},
	}

	for _, testCase := range testCases {
		result := tTwoSumOrderResult(twoSumOrder(testCase.a, testCase.b))
		if len(result) != len(testCase.expected) {
			twoSumPrintError(t, testCase, result)
		}

		for i, n := range testCase.expected {
			if n != result[i] {
				twoSumPrintError(t, testCase, result)
			}
		}
	}
}

func twoSumPrintError(t *testing.T, testCase tTwoSumOrderCase, result tTwoSumOrderResult) {
	t.Fatalf(
		"Case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
