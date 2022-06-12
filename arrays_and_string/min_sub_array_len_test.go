package arrays_and_string

import "testing"

type tMinSubArrayLenResult int

type tMinSubArrayLenCase struct {
	name     string
	a        int
	b        []int
	expected tMinSubArrayLenResult
}

func TestMinSubArrayLen(t *testing.T) {
	testCases := []tMinSubArrayLenCase{
		{
			"Case #1",
			7,
			[]int{2, 3, 1, 2, 4, 3},
			2,
		},
		{
			"Case #2",
			4,
			[]int{1, 4, 4},
			1,
		},
		{
			"Case #3",
			11,
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			0,
		},
		{
			"Case #4",
			15,
			[]int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			2,
		},
	}

	for _, testCase := range testCases {
		result := tMinSubArrayLenResult(minSubArrayLen(testCase.a, testCase.b))
		if result != testCase.expected {
			minSubArrayLenPrintError(t, testCase, result)
		}
	}
}

func minSubArrayLenPrintError(t *testing.T, testCase tMinSubArrayLenCase, result tMinSubArrayLenResult) {
	t.Fatalf(
		"Case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
