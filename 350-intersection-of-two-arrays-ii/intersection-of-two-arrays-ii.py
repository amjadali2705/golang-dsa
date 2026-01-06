class Solution:
    def intersect(self, nums1: List[int], nums2: List[int]) -> List[int]:
        if len(nums1) > len(nums2):
            nums1, nums2 = nums2, nums1
        
        counts = Counter(nums1)
        result = []
        
        for num in nums2:
            if counts[num] > 0:
                result.append(num)
                counts[num] -= 1
                
        return result
        