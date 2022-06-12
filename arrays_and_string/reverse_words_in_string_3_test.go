package arrays_and_string

import "testing"

type tReverseWords3Result string

type tReverseWords3Case struct {
	name     string
	a        string
	expected tReverseWords3Result
}

func TestReverseWords3(t *testing.T) {
	testCases := []tReverseWords3Case{
		{
			"Case #1",
			"Let's take LeetCode contest",
			"s'teL ekat edoCteeL tsetnoc",
		},
		{
			"Case #2",
			"God Ding",
			"doG gniD",
		},
	}

	for _, testCase := range testCases {
		result := tReverseWords3Result(reverseWords3(testCase.a))
		if result != testCase.expected {
			reverseWords3PrintError(t, testCase, result)
		}
	}
}

func reverseWords3PrintError(t *testing.T, testCase tReverseWords3Case, result tReverseWords3Result) {
	t.Fatalf(
		"Case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
