package len_od_last_word

import "testing"

type tResult int

type tCase struct {
	name     string
	input    string
	expected tResult
}

func tPrint(t *testing.T, testCase tCase, result tResult) {
	t.Fatalf("%v. Expected: %v. Actual: %v.", testCase.name, testCase.expected, result)
}

func TestLenOfLastWord(t *testing.T) {
	testCases := []tCase{
		{
			"case #1",
			"Hello World",
			5,
		},
		{
			"case #2",
			"   fly me   to   the moon  ",
			4,
		},
		{
			"case #3",
			"luffy is still joyboy",
			6,
		},
		{
			"case #4",
			"a",
			1,
		},
	}

	for _, testCase := range testCases {
		result := tResult(lenOfLastWord(testCase.input))
		if result != testCase.expected {
			tPrint(t, testCase, result)
		}
	}
}
