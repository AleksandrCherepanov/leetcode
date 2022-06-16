package recursion

import "testing"

type swapPairsTestCase struct {
	name     string
	head     *ListNode
	expected *ListNode
}

type swapPairsResult *ListNode

func TestSwapPairs(t *testing.T) {
	testCases := []swapPairsTestCase{
		{
			"Case #1",
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
			},
		},
		{
			"Case #2",
			nil,
			nil,
		},
		{
			"Case #3",
			&ListNode{
				Val:  1,
				Next: nil,
			},
			&ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	for _, testCase := range testCases {
		result := swapPairsResult(swapPairs(testCase.head))

		for result != nil && testCase.expected != nil {
			if result.Val != testCase.expected.Val {
				swapPairsPrintResult(t, testCase, result)
			}
			result = result.Next
			testCase.expected = testCase.expected.Next
		}
	}
}

func swapPairsPrintResult(t *testing.T, testCase swapPairsTestCase, result swapPairsResult) {
	t.Errorf(`%v: Expected: %v, Actual: %v`, testCase.name, testCase.expected, result)
}
