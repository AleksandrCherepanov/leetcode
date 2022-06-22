package recursion

import (
	"math"
	"testing"
)

type myPowTestCase struct {
	name     string
	x        float64
	n        int
	expected myPowResult
}

type myPowResult float64

func TestMyPow(t *testing.T) {
	testCases := []myPowTestCase{
		{
			"Case #1",
			2.10000,
			3,
			9.26100,
		},
		{
			"Case #2",
			2.00000,
			-2,
			0.25000,
		},
		{
			"Case #3",
			2.00000,
			0,
			1,
		},
		{
			"Case #4",
			8.88023,
			3,
			700.28148,
		},
		{
			"Case #5",
			0.00001,
			int(math.MaxInt32),
			0.00000,
		},
	}

	for _, testCase := range testCases {
		result := myPowResult(myPow(testCase.x, testCase.n))

		if result != testCase.expected {
			myPowPrintResult(t, testCase, result)
		}

	}
}

func myPowPrintResult(t *testing.T, testCase myPowTestCase, result myPowResult) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
