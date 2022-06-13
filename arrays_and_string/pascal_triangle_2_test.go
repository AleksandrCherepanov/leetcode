package arrays_and_string

import "testing"

type tPascalTriangle2Result []int

type tPascalTriangle2Case struct {
	name     string
	input    int
	expected tPascalTriangle2Result
}

func TestPascalTriangle2(t *testing.T) {
	testCases := []tPascalTriangle2Case{
		{
			"simple case",
			0,
			[]int{1},
		},
		{
			"middle case",
			2,
			[]int{1, 2, 1},
		},
		{
			"hard case",
			4,
			[]int{1, 4, 6, 4, 1},
		},
		{
			"mega hard case",
			7,
			[]int{1, 7, 21, 35, 35, 21, 7, 1},
		},
	}

	for _, testPascalTriangle2Case := range testCases {
		result := tPascalTriangle2Result(pascalTriangle2(testPascalTriangle2Case.input))

		if len(result) != len(testPascalTriangle2Case.expected) {
			pascalTriangle2printError(t, testPascalTriangle2Case, result)
		}

		for i, el := range result {
			if el != testPascalTriangle2Case.expected[i] {
				pascalTriangle2printError(t, testPascalTriangle2Case, result)
			}
		}
	}
}

func pascalTriangle2printError(t *testing.T, testPascalTriangle2Case tPascalTriangle2Case, result tPascalTriangle2Result) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testPascalTriangle2Case.name,
		testPascalTriangle2Case.expected,
		result,
	)
}
