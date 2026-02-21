class Solution:
    def shortestToChar(self, s: str, c: str) -> List[int]:
        n = len(s)
        res = [0] * n
        
        pos = float('-inf')
        
        for i in range(n):
            if s[i] == c:
                pos = i
            res[i] = i - pos
            
        pos = float('inf')
        for i in range(n - 1, -1, -1):
            if s[i] == c:
                pos = i
            res[i] = min(res[i], pos - i)
            
        return res