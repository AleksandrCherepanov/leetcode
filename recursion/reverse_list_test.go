package recursion

import (
	"testing"
)

type reverseListTestCase struct {
	name     string
	head     *ListNode
	expected *ListNode
}

type reverseListResult *ListNode

func TestReverseList(t *testing.T) {
	testCases := []reverseListTestCase{
		{
			"Case #1",
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
			&ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  1,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			"Case #2",
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  2,
					Next: nil,
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	for _, testCase := range testCases {
		result := reverseListResult(reverseList(testCase.head))

		res := make([]int, 0, 0)
		exp := make([]int, 0, 0)

		CollectListValues(result, &res)
		CollectListValues(testCase.expected, &exp)

		if len(res) != len(exp) {
			reverseListPrintResult(t, testCase, res, exp)
		}

		for i, n := range res {
			if n != exp[i] {
				reverseListPrintResult(t, testCase, res, exp)
			}
		}
	}
}

func CollectListValues(head *ListNode, values *[]int) {
	if head == nil {
		return
	}

	*values = append(*values, head.Val)
	CollectListValues(head.Next, values)
}

func reverseListPrintResult(t *testing.T, testCase reverseListTestCase, res, exp []int) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, exp, res)
}
