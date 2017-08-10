#!/usr/bin/env python

# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution(object):
    def addTwoNumbers(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        
        result = l1
        while l1 and l2 :
            l1.val += l2.val
            
            if l1.val >= 10:
                l1.val -= 10
                
                try:
                    l1.next.val += 1
                except:
                    l1.next = ListNode(1)
            
            if not l1.next:
                l1.next = l2.next
                break
                
            l1 = l1.next
            l2 = l2.next
        
        while l1:
            if l1.val < 10:
                if not l1.next or l1.next.val < 10:
                    break
            
            l1.val -= 10
            
            try:
                l1.next.val += 1
            except:
                l1.next = ListNode(1)
            l1 = l1.next
        
        return result
            

if __name__ == "__main__":
    x = []
    for i in [2, 4, 3]:
        x.append(ListNode(i))

    y = []
    for i in [5, 6, 4]:
        y.append(ListNode(i))

    for i in range(2):
        x[i].next = x[i+1]
        y[i].next = y[i+1]

    obj = Solution()
    x = obj.addTwoNumbers(x[0], y[0])
        
    z = []
    while x:
        z.append(str(x.val))
        x = x.next

    print " -> ".join(z)

# vim: ts=4 sw=4 expandtab
