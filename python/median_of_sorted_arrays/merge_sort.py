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


        # Merge sort
        i = 0
        j = 0
        nums3 = []
        while i < len(nums1) and j < len(nums2):
            if nums1[i] <= nums2[j]:
                nums3.append(nums1[i])
                i += 1
            elif nums1[i] > nums2[j]:
                nums3.append(nums2[j])
                j += 1 

        if i == len(nums1):
            nums3 += nums2[j:]
        elif j == len(nums2):
            nums3 += nums1[i:]

        # Get median
        if len(nums3)%2 == 1:
            return float(nums3[(len(nums3)-1)/2])

        m = len(nums3)/2
        return (float(nums3[m-1]) + float(nums3[m] - nums3[m-1])/2.0)
        
if __name__ == "__main__":

    obj = Solution()
    print obj.findMedianSortedArrays([1,3], [2])        
    print obj.findMedianSortedArrays([1], [2, 3, 4, 5, 6, 7 ,8])        
    print obj.findMedianSortedArrays(range(55,90,3), range(100,600,7))        
    print obj.findMedianSortedArrays([1,2], [3,4])        

# vim: ts=4 sw=4 expandtab
