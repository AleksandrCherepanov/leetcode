def removeNthFromEnd(head, n):
    if head.next is None:
        return None

    slow = head
    fast = head

    for i in range(0, n):
        fast = fast.next

    while fast is not None and fast.next is not None:
        slow = slow.next
        fast = fast.next

    if fast is None:
        head = slow.next
    else:
        slow.next = slow.next.next

    return head
