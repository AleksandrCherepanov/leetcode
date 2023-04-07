package main

import "math"

type ListNode struct {
    Val int
    Next *ListNode
}

func getDecimalValue(head *ListNode) int {
    head = reverse(head)

    sum := 0
    mlt := 0
    for head != nil {
        sum += head.Val * int(math.Pow(2, float64(mlt)))
        mlt++
        head = head.Next
    }

    return sum
}

func reverse(head *ListNode) *ListNode {
    prev := head
    curr := head.Next

    for curr != nil {
        nxt := curr.Next
        curr.Next = prev
        prev = curr
        curr = nxt
    }

    head.Next = curr
    return prev
}

