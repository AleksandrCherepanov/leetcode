package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head.Next == nil {
		return nil
	}

	slow := head
	fast := head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	if fast == nil {
		head = slow.Next
	} else {
		slow.Next = slow.Next.Next
	}

	return head
}
