package arrays_and_string

import "testing"

type tLongestCommonPrefixResult string

type tLongestCommonPrefixCase struct {
	name     string
	a        []string
	expected tLongestCommonPrefixResult
}

func TestLongestCommonPrefix(t *testing.T) {
	testCases := []tLongestCommonPrefixCase{
		{
			"Case #1",
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			"Case #2",
			[]string{"dog", "racecar", "car"},
			"",
		},
	}

	for _, testCase := range testCases {
		result := tLongestCommonPrefixResult(longestCommonPrefix(testCase.a))
		if result != testCase.expected {
			longestCommonPrefixPrintError(t, testCase, result)
		}
	}
}

func longestCommonPrefixPrintError(t *testing.T, testCase tLongestCommonPrefixCase, result tLongestCommonPrefixResult) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
