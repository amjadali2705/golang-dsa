class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        counts = {}
        
        for num in nums:
            counts[num] = counts.get(num, 0) + 1

        for count in counts.values():
            if count > 1:
                return True
                
        return False
        