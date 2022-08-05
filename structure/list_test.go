package structure

import "testing"

func TestListPush(t *testing.T) {
	list := NewNode(1)
	list.Push(2)

	if list.Next == nil || list.Next.Val != 2 {
		t.Fatalf("List push invalid: Expected: %v. Actual: %v.", 2, list.Next)
	}

	list.Push(3)

	if list.Next.Next == nil || list.Next.Next.Val != 3 {
		t.Fatalf("List push invalid: Expected: %v. Actual: %v.", 3, list.Next.Next)
	}

	if list.Next.Next.Next != nil {
		t.Fatalf("List push invalid: Expected: %v. Actual: %v.", nil, list.Next.Next.Next)
	}
}

func TestListToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		input    *ListNode
		expected []int
	}{
		{
			"Empty list to slice",
			&ListNode{},
			[]int{0},
		},
		{
			"Non empty list to slice",
			NewNode(1).Push(2).Push(3).Push(4).Push(5),
			[]int{1, 2, 3, 4, 5},
		},
	}

	for _, testCase := range testCases {
		actual := testCase.input.ToSlice()
		if len(testCase.expected) != len(actual) {
			t.Fatalf("List to slice invalid: Expected: %v. Actual: %v.", testCase.expected, actual)
		}

		for i, v := range actual {
			if v != testCase.expected[i] {
				t.Fatalf("List to slice invalid: Expected: %v. Actual: %v.", testCase.expected, actual)
			}
		}
	}
}
