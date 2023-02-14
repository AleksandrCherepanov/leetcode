def reverseBetween(head, left, right):
    if head is None or head.next is None:
        return head

    if right - left == 0:
        return head
    
    dummy = new ListNode(0, head)
    leftPtr = dummy

    prev = head
    i = 1 
    while i < left:
        leftPtr = leftPtr.next
        i += 1

    prev = leftPtr
    curr = leftPtr.next
    
    i = left
    while i <= right:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt
        i += 1
    
    leftPtr.next.next = curr
    leftPtr.next = prev

    return dummy.next
