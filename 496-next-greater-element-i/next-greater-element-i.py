class Solution:
    def nextGreaterElement(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = []
        for num in nums1:
            nums2Indx = nums2.index(num)

            found = False
            for j in range(nums2Indx + 1, len(nums2)):
                if nums2[j] > num:
                    res.append(nums2[j])
                    found = True
                    break

            if not found:
                res.append(-1)

        return res
