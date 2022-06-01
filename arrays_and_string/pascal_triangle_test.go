package arrays_and_string

import "testing"

type tResult [][]int

type tCase struct {
	name     string
	input    int
	expected tResult
}

func TestPascalTriangle(t *testing.T) {
	testCases := []tCase{
		{
			"simple case",
			1,
			[][]int{{1}},
		},
		{
			"middle case",
			3,
			[][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			"hard case",
			5,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			"mega hard case",
			8,
			[][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}, {1, 5, 10, 10, 5, 1}, {1, 6, 15, 20, 15, 6, 1}, {1, 7, 21, 35, 35, 21, 7, 1}},
		},
	}

	for _, testCase := range testCases {
		result := tResult(Generate(testCase.input))

		if len(result) != len(testCase.expected) {
			printError(t, testCase, result)
		}

		for i, line := range result {
			if len(testCase.expected[i]) != len(line) {
				printError(t, testCase, result)
			}

			for j, n := range line {
				if line[j] != n {
					printError(t, testCase, result)
				}
			}
		}
	}
}

func printError(t *testing.T, testCase tCase, result tResult) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
