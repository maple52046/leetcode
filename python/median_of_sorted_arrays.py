#!/usr/bin/env python

class Solution(object):
    def findMedianSortedArrays(self, nums1, nums2):
        """
        :type nums1: List[int]
        :type nums2: List[int]
        :rtype: float
        """
        
        # Make the list nums1 always longger than nums2
        if len(nums1) < len(nums2):
            nums1, nums2 = nums2, nums1
        
        # Extented nums1 with element of nums2 by using binary search
        for n in nums2:
            x, y = self.get_median(nums1)
            p = self.binary_search(nums1, x ,n)
            nums1.insert(p, n)

        # Get median
        x, y = self.get_median(nums1)
        if y == -1:
            return float(nums1[x])

        return (float(nums1[x]) + float(nums1[y] - nums1[x])/2.0)


            
    def binary_search(self, array, m, n):
        # Binary search with recursive

        if len(array) == 1:
            return (m+1) if n >= array[0] else m

        if len(array) == 2:
            if n < array[0]:
                return 0
            elif n > array[1]:
                return 2
            return 1

        print "m={}, array_leng: {}".format(m, len(array))
        if n == array[m]:
            return (m+1)

        if n > array[m]:
            x, y = self.get_median(array[m+1:])
            print "next array: {}".format(array[m+1:])
            return self.binary_search(array[m+1:], x, n) + m + 1
        
        x, y = self.get_median(array[:m])
        print "next array: {}".format(array[:m])
        return self.binary_search(array[:m], x, n)
            
    
    def get_median(self, array):
        if len(array) <= 1:
            return 0, -1
        
        if len(array)%2 == 0:
            m = len(array)/2
            return (m-1, m)
        
        return (len(array)-1)/2, -1


if __name__ == "__main__":

    obj = Solution()
    print obj.findMedianSortedArrays([1,3], [2])        
    print obj.findMedianSortedArrays([1], [2, 3, 4, 5, 6, 7 ,8])        
    print obj.findMedianSortedArrays(range(55,90,3), range(100,600,7))        

# vim: ts=4 sw=4 expandtab
