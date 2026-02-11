class Solution:
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        res = []
        minSum = float(inf)

        for i in range(len(list1)):
            if list1[i] in list2:
                j = list2.index(list1[i])

                if i + j < minSum:
                    minSum = i + j
                    res = [list1[i]]
                elif i + j == minSum:
                    res.append(list1[i])

        return res