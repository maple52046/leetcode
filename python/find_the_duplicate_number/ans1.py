class Solution:
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        max_num = nums[0]
        counts = {}

        for n in nums:
            try:
                counts[n] += 1
            except KeyError:
                counts[n] = 1
            if n != max_num:
                if counts[n] > counts[max_num]:
                    max_num = n
        return max_num

if __name__ == "__main__":
    s = Solution()
    print(s.findDuplicate([1,3,4,2,2]))
    print(s.findDuplicate([3,1,3,4,2]))
