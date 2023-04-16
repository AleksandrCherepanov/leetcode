package main

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	odd := head
	even := odd.Next

	for even != nil && even.Next != nil {
		oNext := odd.Next
		odd.Next = even.Next
		even.Next = even.Next.Next
		odd.Next.Next = oNext

		odd = odd.Next
		even = even.Next
	}

	return head
}
