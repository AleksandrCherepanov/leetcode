package main

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	middle := head.Next
	prev := &ListNode{0, nil}

	for head != nil && head.Next != nil {
		if prev != nil {
			prev.Next = head.Next
		}
		prev = head
		nextNode := head.Next.Next
		head.Next.Next = head

		head.Next = nextNode
		head = nextNode
	}

	return middle
}

func swapPairsRecursive(head *ListNode) *ListNode {
	newHead := swap(head)
	return newHead
}

func swap(node *ListNode) *ListNode {
	if node == nil || node.Next == nil {
		return node
	}

	second := node.Next
	third := second.Next

	second.Next = node
	node.Next = swap(third)

	return second
}
