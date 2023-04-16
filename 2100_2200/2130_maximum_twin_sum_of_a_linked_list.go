package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//TODO need to try rewrite it in recursive way

func pairSum(head *ListNode) int {
	length := lenList(head)
	head = reverse(head)

	mid := length / 2
	second := head

	for i := 0; i < mid; i++ {
		second = second.Next
	}

	second = reverse(second)

	maxSum := 0
	curr := head
	for second != nil {
		sum := curr.Val + second.Val
		if sum > maxSum {
			maxSum = sum
		}
		curr = curr.Next
		second = second.Next
	}

	return maxSum
}

func lenList(head *ListNode) int {
	l := 0
	curr := head
	for curr != nil {
		curr = curr.Next
		l++
	}

	return l
}

func reverse(head *ListNode) *ListNode {
	prev := head
	curr := prev.Next

	nxt := &ListNode{0, nil}
	for curr != nil {
		nxt = curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	head.Next = curr
	return prev
}
