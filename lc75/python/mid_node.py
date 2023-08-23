# Definition for singly-linked list.
from typing import List, Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        
    def __str__(self):
        s = str(self.val)
        if self.next is not None:
            s += ", " + str(self.next)
        return s


class Solution:
    # get to seeker=6, mid=7. 6 has no next, return 7
    def deleteMiddle(self, head: Optional[ListNode]) -> Optional[ListNode]:
        seeker = head
        mid = head
        last_mid = head
        i = 0
        while seeker is not None:
            seeker = seeker.next
            i += 1
            if seeker is not None:
                seeker = seeker.next
                i += 1
                last_mid = mid
                mid = mid.next
            else:
                break
        if i == 1:
            return None
        last_mid.next = mid.next
        return head
    
    
def to_linked_list(l: List[int]) -> Optional[ListNode]:
    if len(l) == 0:
        return None
    ll = ListNode(l[0])
    ll.next = to_linked_list(l[1:])
    return ll

s = Solution()
print(s.deleteMiddle(to_linked_list([1])))


