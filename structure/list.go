package structure

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func (first *ListNode) Push(val int) *ListNode {
	iter := first
	for iter.Next != nil {
		iter = iter.Next
	}

	iter.Next = NewNode(val)

	return first
}

func (ln *ListNode) IsEqual(listNode *ListNode) bool {
	head1 := ln
	head2 := listNode
	for head1 != nil && head2 != nil {
		if head1.Val != head2.Val {
			return false
		}
		head1 = head1.Next
		head2 = head2.Next
	}

	if head1 != nil || head2 != nil {
		return false
	}

	return true
}

func (first *ListNode) ToSlice() []int {
	iter := first
	result := make([]int, 0, 0)
	for iter != nil {
		result = append(result, iter.Val)
		iter = iter.Next
	}

	return result
}
