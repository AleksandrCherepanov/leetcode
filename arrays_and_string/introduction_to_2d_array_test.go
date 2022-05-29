package arrays_and_string

import "testing"

func TestDiagonalTraverse(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			"empty case",
			[][]int{},
			[]int{},
		},
		{
			"case #1",
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			"case #2",
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 3, 4},
		},
		{
			"case #3",
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[]int{1, 2, 3, 5, 4, 6},
		},
		{
			"case #4",
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}},
			[]int{1, 2, 5, 6, 3, 4, 7, 8},
		},
	}

	for _, testCase := range testCases {
		result := DiagonalTraverse(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Errorf(
				"Fail. Test case: %s. Expected len: %v. Actual len: %v",
				testCase.name,
				len(testCase.expected),
				len(result),
			)
		}

		for i, n := range result {
			if testCase.expected[i] != n {
				t.Errorf(
					"Fail. Test case: %s. Expected: %v. Actual: %v",
					testCase.name,
					testCase.expected,
					result,
				)
				break
			}
		}
	}
}
