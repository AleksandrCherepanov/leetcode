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

func TestRemoveElement(t *testing.T) {
	var testCases = []struct {
		caseName string
		input    []int
		value    int
		expected int
	}{
		{
			"empty array",
			[]int{},
			4,
			0,
		},
		{
			"common case #1",
			[]int{3, 2, 2, 3},
			3,
			2,
		},
		{
			"common case #2",
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
		},
		{
			"no value in array",
			[]int{3, 2, 2, 3},
			4,
			4,
		},
		{
			"all elements",
			[]int{2, 2, 2, 2},
			2,
			0,
		},
	}
	for _, testCase := range testCases {
		result := removeElement(testCase.input, testCase.value)

		if result != testCase.expected {
			t.Fatalf(
				"Arrays are not equal. Test case: %s. Expected: %v. Actual: %v",
				testCase.caseName,
				testCase.expected,
				result,
			)
		}
	}
}

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
