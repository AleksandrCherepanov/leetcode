package recursion

import (
	"fmt"
	"testing"
)

type mergeTwoSortedListsTestCase struct {
	name     string
	l1       *ListNode
	l2       *ListNode
	expected *ListNode
}

type mergeTwoSortedListsResult *ListNode

func TestMergeTwoSortedLists(t *testing.T) {
	testCases := []mergeTwoSortedListsTestCase{
		{
			"Case #1",
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  4,
						Next: nil,
					},
				},
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			"Case #2",
			nil,
			nil,
			nil,
		},
		{
			"Case #3",
			nil,
			&ListNode{
				Val:  0,
				Next: nil,
			},
			&ListNode{
				Val:  0,
				Next: nil,
			},
		},
	}

	for _, testCase := range testCases {
		result := mergeTwoSortedListsResult(mergeTwoSortedLists(testCase.l1, testCase.l2))
		fmt.Println(testCase.l1)
		res := make([]int, 0, 0)
		exp := make([]int, 0, 0)

		CollectListValues(result, &res)
		CollectListValues(testCase.expected, &exp)

		if len(res) != len(exp) {
			mergeTwoSortedListsPrintResult(t, testCase, res, exp)
		}

		for i, n := range res {
			if n != exp[i] {
				mergeTwoSortedListsPrintResult(t, testCase, res, exp)
			}
		}
	}
}

func mergeTwoSortedListsPrintResult(t *testing.T, testCase mergeTwoSortedListsTestCase, res, exp []int) {
	t.Fatalf(`%v: Expected: %v, Actual: %v`, testCase.name, exp, res)
}
