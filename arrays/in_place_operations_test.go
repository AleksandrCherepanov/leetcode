package arrays

import (
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	var testCases = []struct {
		caseName string
		input    []int
		expected []int
	}{
		{
			"empty array",
			[]int{},
			[]int{},
		},
		{
			"common case",
			[]int{3, 1, 2, 4},
			[]int{2, 4, 3, 1},
		},
		{
			"only even case",
			[]int{2, 4, 6, 8},
			[]int{2, 4, 6, 8},
		},
		{
			"only odd case",
			[]int{1, 3, 5, 7},
			[]int{1, 3, 5, 7},
		},
		{
			"mixed numbers case",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{2, 4, 6, 8, 5, 3, 7, 1, 9},
		},
		{
			"mixed numers start with even",
			[]int{2, 4, 1, 3, 6, 7, 8, 9, 4},
			[]int{2, 4, 6, 8, 4, 7, 3, 9, 1},
		},
	}
	for _, testCase := range testCases {
		result := SortArrayByParity(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Fatalf(
				"Len of arrays is not equal.Test case: %s. Expected: %d. Actual: %d",
				testCase.caseName,
				len(testCase.expected),
				len(result),
			)
		}

		for i := 0; i < len(result); i++ {
			if result[i] != testCase.expected[i] {
				t.Fatalf(
					"Arrays are not equal. Test case: %s. Expected: %v. Actual: %v",
					testCase.caseName,
					testCase.expected,
					result,
				)
			}
		}
	}
}
