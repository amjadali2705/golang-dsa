class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        ranges = []
        if not nums:
            return ranges
        
        start = nums[0]
        
        for i in range(len(nums)):
            if i + 1 == len(nums) or nums[i+1] != nums[i] + 1:
                if nums[i] == start:
                    ranges.append(str(start))
                else:
                    ranges.append(f"{start}->{nums[i]}")
                
                if i + 1 < len(nums):
                    start = nums[i+1]

        return ranges