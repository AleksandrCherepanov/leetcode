package arrays_and_string

import "testing"

type tAddBinaryResult string

type tAddBinaryCase struct {
	name     string
	a        string
	b        string
	expected tAddBinaryResult
}

func TestAddBinary(t *testing.T) {
	testCases := []tAddBinaryCase{
		{
			"simplest case",
			"0",
			"0",
			"0",
		},
		{
			"simple case",
			"0",
			"1",
			"1",
		},
		{
			"middle case",
			"1",
			"1",
			"10",
		},
		{
			"hard case",
			"11",
			"1",
			"100",
		},
		{
			"mega hard case",
			"1010",
			"1011",
			"10101",
		},
	}

	for _, testCase := range testCases {
		result := tAddBinaryResult(addBinary(testCase.a, testCase.b))
		if result != testCase.expected {
			addBinaryPrintError(t, testCase, result)
		}
	}
}

func addBinaryPrintError(t *testing.T, testCase tAddBinaryCase, result tAddBinaryResult) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
