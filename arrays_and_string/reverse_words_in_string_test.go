package arrays_and_string

import "testing"

type tReverseWordsInStringResult string

type tReverseWordsInStringCase struct {
	name     string
	a        string
	expected tReverseWordsInStringResult
}

func TestReverseWordsInString(t *testing.T) {
	testCases := []tReverseWordsInStringCase{
		{
			"Case #1",
			"the sky is blue",
			"blue is sky the",
		},
		{
			"Case #2",
			"  hello world  ",
			"world hello",
		},
		{
			"Case #3",
			"a good   example",
			"example good a",
		},
	}

	for _, testCase := range testCases {
		result := tReverseWordsInStringResult(reverseWordsInString(testCase.a))
		if result != testCase.expected {
			reverseWordsInStringPrintError(t, testCase, result)
		}
	}
}

func reverseWordsInStringPrintError(t *testing.T, testCase tReverseWordsInStringCase, result tReverseWordsInStringResult) {
	t.Fatalf(
		"Case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
