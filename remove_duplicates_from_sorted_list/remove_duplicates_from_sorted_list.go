package remove_duplicates_from_sorted_list

import (
	"github.com/AleksandrCherepanov/leetcode/structure"
)

func RemoveDuplicatesFromSortedList(head *structure.ListNode) *structure.ListNode {
	if head == nil {
		return head
	}

	pointer := head
	prev := pointer
	for pointer != nil {
		if pointer.Val != prev.Val {
			head.Next = pointer
			prev = pointer
		}

		pointer = pointer.Next
	}

	head.Next = pointer

	return head
}
