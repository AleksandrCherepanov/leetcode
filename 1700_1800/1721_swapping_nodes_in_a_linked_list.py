def swapNodes(head, k):
    if head is None:
        return head
    
    if head.next is None:
        return head

    l = 0
    curr = head
    while curr is not None:
        curr = curr.next
        l += 1

    last = l - k

    left = head
    right = head

    for i in range(0, k - 1):
        left = left.next
    
    for i in range(0, last):
        right = right.next
    
    tmp = left.val
    left.val = right.val
    right.val = tmp

    return head
