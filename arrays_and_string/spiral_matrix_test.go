package arrays_and_string

import "testing"

func TestSpiralOrder(t *testing.T) {
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
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			"case #2",
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			"case #3",
			[][]int{{1, 2}, {3, 4}, {5, 6}},
			[]int{1, 2, 4, 6, 5, 3},
		},
		{
			"case #4",
			[][]int{{1, 2}, {3, 4}},
			[]int{1, 2, 4, 3},
		},
	}

	for _, testCase := range testCases {
		result := SpiralOrder(testCase.input)

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
