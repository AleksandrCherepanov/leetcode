package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	dummy := &ListNode{0, head}
	prev := dummy
	curr := head

	for curr != nil {
		if curr.Val == val {
			prev.Next = curr.Next
		} else {
			prev = curr
		}

		curr = curr.Next
	}

	return dummy.Next
}
