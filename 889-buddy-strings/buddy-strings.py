class Solution:
    def buddyStrings(self, s: str, goal: str) -> bool:
        if len(s) != len(goal):
            return False
        
        if s == goal:
            return len(set(s)) < len(s)
        
        diffs = []
        for i in range(len(s)):
            if s[i] != goal[i]:
                diffs.append(i)
                
        if len(diffs) != 2:
            return False
        
        i, j = diffs
        return s[i] == goal[j] and s[j] == goal[i]