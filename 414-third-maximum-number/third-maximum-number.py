class Solution:
    def thirdMax(self, nums: List[int]) -> int:
        first = second = third = float('-inf')
        seen = set()
        count = 0

        for num in nums:
            if num in seen:
                continue
            
            seen.add(num)
            count += 1
            
            if num > first:
                third = second
                second = first
                first = num
            elif num > second:
                third = second
                second = num
            elif num > third:
                third = num
                
        return third if count >= 3 else first