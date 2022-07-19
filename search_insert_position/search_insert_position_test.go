package search_insert_position

import "testing"

type searchInsertPositionResult int

type searchInsertPositionTest struct {
	name     string
	input    []int
	target   int
	expected searchInsertPositionResult
}

func searchInsertPositionPrint(t *testing.T, test searchInsertPositionTest, result searchInsertPositionResult) {
	t.Fatalf("%v. Expected: %v. Actual: %v.", test.name, test.expected, result)
}

func TestSearchInsertPosition(t *testing.T) {
	testCases := []searchInsertPositionTest{
		{
			"case #1",
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			"case #2",
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			"case #3",
			[]int{1, 3, 5, 6},
			7,
			4,
		},
	}

	for _, testCase := range testCases {
		result := searchInsertPositionResult(searchInsertPosition(testCase.input, testCase.target))
		if result != testCase.expected {
			searchInsertPositionPrint(t, testCase, result)
		}
	}
}
