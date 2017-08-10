#!/usr/bin/env python

class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for i in range(len(nums)):
            x = target - nums[i]
            try:
                j = nums.index(x)
                if i != j:
                    return [i, j]
            except:
                pass

if __name__ == "__main__":

    from random import randint

    obj = Solution()
    print obj.twoSum([2,7,11,15], 9)

# vim: ts=4 sw=4 expandtab
