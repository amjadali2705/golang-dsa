class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        numsMap = {}

        for i in range(len(nums)):
            if target - nums[i] in numsMap:
                return [numsMap[target - nums[i]], i]
            numsMap[nums[i]] = i
        
        