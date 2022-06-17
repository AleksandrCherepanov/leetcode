package recursion

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
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
