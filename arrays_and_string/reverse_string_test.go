package arrays_and_string

import "testing"

type tReverseStringResult []byte

type tReverseStringCase struct {
	name     string
	a        []byte
	expected tReverseStringResult
}

func TestReverseString(t *testing.T) {
	testCases := []tReverseStringCase{
		{
			"simplest case",
			[]byte("hello"),
			[]byte("olleh"),
		},
		{
			"simple case",
			[]byte("Hannah"),
			[]byte("hannaH"),
		},
	}

	for _, testCase := range testCases {
		reverseString(testCase.a)
		if len(testCase.a) != len(testCase.expected) {
			reverseStringPrintError(t, testCase, testCase.a)
		}
		for i, j := range testCase.a {
			if j != testCase.expected[i] {
				reverseStringPrintError(t, testCase, testCase.a)
			}
		}
	}
}

func reverseStringPrintError(t *testing.T, testCase tReverseStringCase, result tReverseStringResult) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
