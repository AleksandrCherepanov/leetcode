package remove_duplicates_from_sorted_list

import (
	"github.com/AleksandrCherepanov/leetcode/structure"
)

func RemoveDuplicatesFromSortedList(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return head
	}

	pointer := head
	prev := head
	for pointer != nil {
		if pointer.Val != prev.Val {
			prev.Next = pointer
			prev = pointer
		}

		if pointer.Next == nil && pointer.Val == prev.Val {
			prev.Next = nil
		}

		pointer = pointer.Next
	}

	return head
}
