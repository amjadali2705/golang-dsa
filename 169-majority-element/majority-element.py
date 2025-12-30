class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        counts = {}
        
        for num in nums:
            counts[num] = counts.get(num, 0) + 1
            
        majority_threshold = len(nums) // 2
        
        for num, count in counts.items():
            if count > majority_threshold:
                return num
                
        return -1