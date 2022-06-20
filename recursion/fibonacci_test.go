package recursion

import (
	"testing"
)

type fibonacciTestCase struct {
	name     string
	input    int
	expected fibonacciResult
}

type fibonacciResult int

func TestFibonacci(t *testing.T) {
	testCases := []fibonacciTestCase{
		{
			"Case #1",
			2,
			1,
		},
		{
			"Case #2",
			3,
			2,
		},
		{
			"Case #3",
			4,
			3,
		},
		{
			"Case #4",
			5,
			5,
		},
		{
			"Case #5",
			6,
			8,
		},
		{
			"Case #6",
			15,
			610,
		},
		{
			"Case #7",
			30,
			832040,
		},
	}

	for _, testCase := range testCases {
		result := fibonacciResult(fibonacci(testCase.input))

		if result != testCase.expected {
			fibonacciPrintResult(t, testCase, result)
		}

	}
}

func fibonacciPrintResult(t *testing.T, testCase fibonacciTestCase, result fibonacciResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
