package recursion

import (
	"testing"
)

type kthGrammarTestCase struct {
	name     string
	n        int
	k        int
	expected kthGrammarResult
}

type kthGrammarResult int

func TestKthGrammar(t *testing.T) {
	testCases := []kthGrammarTestCase{
		{
			"Case #1",
			1,
			1,
			0,
		},
		{
			"Case #2",
			2,
			1,
			0,
		},
		{
			"Case #3",
			2,
			2,
			1,
		},
		{
			"Case #4",
			5,
			12,
			1,
		},
	}

	for _, testCase := range testCases {
		result := kthGrammarResult(kthGrammar(testCase.n, testCase.k))

		if result != testCase.expected {
			kthGrammarPrintResult(t, testCase, result)
		}

	}
}

func kthGrammarPrintResult(t *testing.T, testCase kthGrammarTestCase, result kthGrammarResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
