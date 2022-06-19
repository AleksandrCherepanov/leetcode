package recursion

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
