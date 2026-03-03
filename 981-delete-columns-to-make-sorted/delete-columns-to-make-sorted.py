class Solution:
    def minDeletionSize(self, strs: List[str]) -> int:
        delete = 0
        rows = len(strs)
        cols = len(strs[0])
        
        for col in range(cols):
            for row in range(rows - 1):
                if strs[row][col] > strs[row + 1][col]:
                    delete += 1
                    break
                    
        return delete