package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	var nextHead *ListNode

	for i := 1; i < 1e5 && head != nil; i++ {
		tail := head

		j := 1
		for j < i && tail != nil && tail.Next != nil {
			tail = tail.Next
			j++
		}

		nextHead = tail.Next

		if j%2 == 0 {
			tail.Next = nil
			prev.Next = reverse(head)
			prev = head
			head.Next = nextHead
			head = nextHead
		} else {
			prev = tail
			head = nextHead
		}
	}

	return dummy.Next
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	prev := &ListNode{0, nil}

	for head != nil {
		nxt := head.Next
		head.Next = prev
		prev = head
		head = nxt
	}

	return prev
}
