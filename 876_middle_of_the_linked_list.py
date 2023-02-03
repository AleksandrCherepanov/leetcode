class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next    

def middleNode(head):
    slow = head
    fast = head

    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next

    return slow

head = ListNode(1)
two = ListNode(2)
three = ListNode(3)
four = ListNode(4)
five = ListNode(5)

head.next = two
head.next.next = three
head.next.next.next = four
head.next.next.next.next = five

result = middleNode(head)
if result.val != 3:
    print("Error", result.val, 3)
else:
    print("OK")

head = ListNode(1)
two = ListNode(2)
three = ListNode(3)
four = ListNode(4)
five = ListNode(5)
six = ListNode(6)

head.next = two
head.next.next = three
head.next.next.next = four
head.next.next.next.next = five
head.next.next.next.next.next = six

result = middleNode(head)
if result.val != 4:
    print("Error", result.val, 3)
else:
    print("OK")
