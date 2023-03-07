def isPalindrome(self, head: Optional[ListNode]) -> bool:
    curr = head
    length = 0

    if head is None or head.next is None:
        return True

    while curr is not None:
        curr = curr.next
        length += 1

    mid = int(length / 2)
    if length % 2 != 0:
        mid += 1

    curr = head
    prev = head
    for i in range(0, mid):
        prev = curr
        curr = curr.next

    left = prev
    prev = curr
    curr = curr.next

    while curr is not None:
        nxt = curr.next
        curr.next = prev
        prev = curr
        curr = nxt

    left.next.next = None
    left.next = prev

    print(left.val)
 
    right = left.next
    left = head

    while right is not None:
        if left.val != right.val:
            return False

        left = left.next
        right = right.next

    return True
