package arrays_and_string

import "testing"

func TestFindPivotIndex(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"empty case",
			[]int{},
			-1,
		},
		{
			"case #1",
			[]int{1, 7, 3, 6, 5, 6},
			3,
		},
		{
			"case #2",
			[]int{1, 2, 3},
			-1,
		},
		{
			"case #3",
			[]int{2, 1, -1},
			0,
		},
	}

	for _, testCase := range testCases {
		result := pivotIndex(testCase.input)

		if result != testCase.expected {
			t.Errorf(
				"Fail. Test case: %s. Expected: %v. Actual: %v",
				testCase.name,
				testCase.expected,
				result,
			)
		}
	}
}

func TestDominantIndex(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			"empty case",
			[]int{},
			-1,
		},
		{
			"case #1",
			[]int{3, 6, 1, 0},
			1,
		},
		{
			"case #2",
			[]int{1, 2, 3, 4},
			-1,
		},
		{
			"case #3",
			[]int{1},
			0,
		},
		{
			"case #4",
			[]int{0, 0, 0, 1},
			3,
		},
		{
			"case #5",
			[]int{3, 0, 0, 2},
			-1,
		},
	}

	for _, testCase := range testCases {
		result := dominantIndex(testCase.input)

		if result != testCase.expected {
			t.Errorf(
				"Fail. Test case: %s. Expected: %v. Actual: %v",
				testCase.name,
				testCase.expected,
				result,
			)
		}
	}
}

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			"empty case",
			[]int{},
			[]int{},
		},
		{
			"case #1",
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			"case #2",
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},
		{
			"case #3",
			[]int{9},
			[]int{1, 0},
		},
		{
			"case #4",
			[]int{9, 9, 9, 9},
			[]int{1, 0, 0, 0, 0},
		},
	}

	for _, testCase := range testCases {
		result := plusOne(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Errorf(
				"Fail. Test case: %s. Expected: %d. Actual: %d",
				testCase.name,
				len(testCase.expected),
				len(result),
			)
		}

		for i := 0; i < len(result); i++ {
			if result[i] != testCase.expected[i] {
				t.Errorf(
					"Fail. Test case: %s. Expected: %v. Actual: %v",
					testCase.name,
					testCase.expected,
					result,
				)
			}
		}
	}
}

func TestSpiralOrder(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
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
	}

	for _, testCase := range testCases {
		result := spiralOrder(testCase.input)

		if len(result) != len(testCase.expected) {
			t.Errorf(
				"Fail. Test case: %s. Expected: %d. Actual: %d",
				testCase.name,
				len(testCase.expected),
				len(result),
			)
		}

		for i := 0; i < len(result); i++ {
			if result[i] != testCase.expected[i] {
				t.Errorf(
					"Fail. Test case: %s. Expected: %v. Actual: %v",
					testCase.name,
					testCase.expected,
					result,
				)
			}
		}
	}
}
