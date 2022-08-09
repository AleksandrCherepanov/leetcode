package remove_duplicates_from_sorted_list

import (
	"testing"

	"github.com/AleksandrCherepanov/leetcode/structure"
)

type tCase struct {
	name     string
	head     *structure.ListNode
	expected *structure.ListNode
}

func TestRemoveDuplicatesFromSortedList(t *testing.T) {
	testCases := []tCase{
		{
			"case #1",
			structure.NewNode(1).Push(1).Push(2),
			structure.NewNode(1).Push(2),
		},
		{
			"case #2",
			structure.NewNode(1).Push(1).Push(2).Push(3).Push(3),
			structure.NewNode(1).Push(2).Push(3),
		},
	}

	for _, testCase := range testCases {
		result := RemoveDuplicatesFromSortedList(testCase.head)

		if !testCase.expected.IsEqual(result) {
			t.Errorf("%v. Expected: %v. Actual: %v.", testCase.name, testCase.expected.ToSlice(), result.ToSlice())
		}
	}
}
