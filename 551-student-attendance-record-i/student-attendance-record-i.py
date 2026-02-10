class Solution:
    def checkRecord(self, s: str) -> bool:
        absentCount = 0
        lateCount = 0
        for i in range(len(s)):
            if s[i] == "A":
                absentCount += 1
                if absentCount == 2:
                    return False

            if s[i] == "L":
                lateCount += 1
                if lateCount == 3:
                    return False
            else:
                lateCount = 0

        return True