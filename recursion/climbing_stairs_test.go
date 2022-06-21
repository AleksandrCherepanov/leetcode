package recursion

import (
	"testing"
)

type climbingStairsTestCase struct {
	name     string
	input    int
	expected climbingStairsResult
}

type climbingStairsResult int

func TestClimbingStairs(t *testing.T) {
	testCases := []climbingStairsTestCase{
		{
			"Case #1",
			2,
			2,
		},
		{
			"Case #2",
			3,
			3,
		},
		{
			"Case #3",
			4,
			5,
		},
		{
			"Case #4",
			5,
			8,
		},
		{
			"Case #5",
			6,
			13,
		},
	}

	for _, testCase := range testCases {
		result := climbingStairsResult(climbingStairs(testCase.input))

		if result != testCase.expected {
			climbingStairsPrintResult(t, testCase, result)
		}

	}
}

func climbingStairsPrintResult(t *testing.T, testCase climbingStairsTestCase, result climbingStairsResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
