package arrays_and_string

import "testing"

type tStrStrResult int

type tStrStrCase struct {
	name     string
	a        string
	b        string
	expected tStrStrResult
}

func TestStrStr(t *testing.T) {
	testCases := []tStrStrCase{
		{
			"Case #1",
			"",
			"test",
			-1,
		},
		{
			"Case #2",
			"test",
			"",
			0,
		},
		{
			"Case #3",
			"hello",
			"ll",
			2,
		},
		{
			"Case #4",
			"aaaaa",
			"bba",
			-1,
		},
		{
			"Case #5",
			"mississippi",
			"issip",
			4,
		},
	}

	for _, testCase := range testCases {
		result := tStrStrResult(strStr(testCase.a, testCase.b))
		if result != testCase.expected {
			strStrPrintError(t, testCase, result)
		}
	}
}

func strStrPrintError(t *testing.T, testCase tStrStrCase, result tStrStrResult) {
	t.Fatalf(
		"Fail. Test case: %s. Expected: %v. Actual: %v",
		testCase.name,
		testCase.expected,
		result,
	)
}
