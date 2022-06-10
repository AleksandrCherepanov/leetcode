package arrays_and_string

import "testing"

type tArrayPartition1 int

type tArrayPartition1Case struct {
	name     string
	input    []int
	expected tArrayPartition1
}

func TestArrayPartition1(t *testing.T) {
	testCases := []tArrayPartition1Case{
		{
			"Case #1",
			[]int{},
			0,
		},
		{
			"Case #2",
			[]int{1, 4, 3, 2},
			4,
		},
		{
			"Case #3",
			[]int{6, 2, 6, 5, 1, 2},
			9,
		},
	}

	for _, testCase := range testCases {
		result := tArrayPartition1(arrayPartition1(testCase.input))

		if result != testCase.expected {
			arrayPartition1PrintError(t, testCase, result)
		}
	}
}

func arrayPartition1PrintError(t *testing.T, testCase tArrayPartition1Case, result tArrayPartition1) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
