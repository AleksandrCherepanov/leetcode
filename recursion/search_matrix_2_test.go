package recursion

import "testing"

type searchMatrix2Result bool

type searchMatrix2TestCase struct {
	name     string
	matrix   [][]int
	target   int
	expected searchMatrix2Result
}

func searchMatrix2PrintError(t *testing.T, testCase searchMatrix2TestCase, result searchMatrix2Result) {
	t.Fatalf("%v. Expected: %v. Actual: %v.", testCase.name, testCase.expected, result)
}

func TestSearchMatrix2(t *testing.T) {
	testCases := []searchMatrix2TestCase{
		{
			"case #1",
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			5,
			true,
		},
		{
			"case #2",
			[][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}},
			20,
			false,
		},
	}

	for _, testCase := range testCases {
		result := searchMatrix2Result(searchMatrix2(testCase.matrix, testCase.target))

		if result != testCase.expected {
			searchMatrix2PrintError(t, testCase, result)
		}

	}
}
