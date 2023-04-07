def deleteDuplicates(head):
    slow = head
    fast = head

    while fast.next:
        fast = fast.next
        if fast.next.val != slow.val:
            slow.next = fast 
            slow = slow.next
    
    if fast.val != slow.val:
        slow.next = fast 
        slow = slow.next

    return head

