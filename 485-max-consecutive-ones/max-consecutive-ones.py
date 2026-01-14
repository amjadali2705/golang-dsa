class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        count = 0
        maxCount = 0

        for num in nums:
            if num == 0:
                count = 0
            else:
                count += 1

            if maxCount < count:
                maxCount = count

        return maxCount
        