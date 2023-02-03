def swapPairs(head):
    if head is None or head.next is None:
        return head
    
    middle = head.next
    prev = None
    
    while head and head.next: 
        if prev is not None:
            prev.next = head.next
        prev = head
        
        nextNode = head.next.next
        head.next.next = head

        head.next = nextNode
        head = nextNode

    return middle

def swap (head):
    if head is None or head.next is None:
		return head

	second = head.next
	third = second.next

	second.next = head
	head.next = swap(third)

	return second

def swapPairsRecursive(head):
    return swap(head)
