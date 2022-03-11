package arrays

import "testing"

func TestHeightChecker(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"empty array",
			[]int{},
			0,
		},
		{
			"common case #1",
			[]int{1, 1, 4, 2, 1, 3},
			3,
		},
		{
			"common case #2",
			[]int{5, 1, 2, 3, 4},
			5,
		},
		{
			"common case #3",
			[]int{1, 2, 3, 4, 5},
			0,
		},
	}

	for _, testCase := range testCases {
		result := heightCheck(testCase.input)

		if result != testCase.expected {
			t.Fatalf(
				"Fail. Test case: %s. Expected: %v. Actual: %v.",
				testCase.name,
				testCase.expected,
				result,
			)
		}
	}
}

func TestThirdMax(t *testing.T) {
	var testCases = []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"empty array",
			[]int{},
			0,
		},
		{
			"common case #1",
			[]int{1},
			1,
		},
		{
			"common case #2",
			[]int{1, 2},
			2,
		},
		{
			"common case #3",
			[]int{3, 2, 1},
			1,
		},
		{
			"common case #4",
			[]int{2, 2, 3, 1},
			1,
		},
		{
			"common case #5",
			[]int{2, 2, 3, 1, 5, 5, 4},
			3,
		},
		{
			"common case #6",
			[]int{1, 2, -2147483648},
			-2147483648,
		},
		{
			"common case #7",
			[]int{1, 1, 2},
			2,
		},
	}

	for _, testCase := range testCases {
		result := thirdMax(testCase.input)

		if result != testCase.expected {
			t.Fatalf(
				"Fail. Test case: %s. Expected: %v. Actual: %v.",
				testCase.name,
				testCase.expected,
				result,
			)
		}
	}
}

func TestFindDisappearedNumbers(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"Empty case",
			[]int{},
			[]int{},
		},
		{
			"Case #1",
			[]int{4, 3, 2, 7, 8, 2, 3, 1},
			[]int{5, 6},
		},
		{
			"Case #2",
			[]int{1, 1},
			[]int{2},
		},
	}

	for _, testCase := range testCases {
		result := findDisappearedNumbers(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Fatalf(
				"Case: %s. Length is not the same. Expected: %v. Actual: %v",
				testCase.name,
				len(testCase.expected),
				len(result),
			)
		}

		for i, v := range result {
			if testCase.expected[i] != v {
				t.Fatalf(
					"Case: %s. Result is not the same. Expected: %v. Actual: %v",
					testCase.name,
					testCase.expected,
					result,
				)
			}
		}
	}
}
