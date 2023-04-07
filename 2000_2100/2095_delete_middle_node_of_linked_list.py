def deleteMiddle(head):
    if head.next is None:
        return head.next

    prev = None
    slow = head
    fast = head

    while fast is not None and fast.next is not None:
        fast = fast.next.next
        prev = slow
        slow = slow.next

    prev.next = slow.next

    return head

