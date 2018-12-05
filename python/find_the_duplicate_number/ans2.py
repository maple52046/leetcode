class Solution:
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        for i in range(len(nums)):
            j = abs(nums[i])
            if nums[j] < 0:
                return j
            nums[j] *= -1

if __name__ == "__main__":
    s = Solution()
    print(s.findDuplicate([1,3,4,2,2]))
    print(s.findDuplicate([3,1,3,4,2]))

