class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        nonZeroPtr = 0
        for i in range(len(nums)):
            if nums[i] != 0:
                nums[nonZeroPtr] = nums[i]
                nonZeroPtr += 1

        for i in range(nonZeroPtr, len(nums)):
            nums[i] = 0
        