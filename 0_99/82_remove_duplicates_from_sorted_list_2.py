
def deleteDuplicates(head):
    if head is None or head.next is None:
        return head
    
    dummy = ListNode(0, head)
    prev = dummy
    start = None
    curr = head
    
    while curr.next is not None:
        if curr.next.val == curr.val:
            start = prev
        elif start is not None:
            start.next = curr.next
            prev = start
            start = None
        else:
            prev = curr
        
        curr = curr.next
    
    if start is not None:
        start.next = curr.next
        start = None
    
    return dummy.next
