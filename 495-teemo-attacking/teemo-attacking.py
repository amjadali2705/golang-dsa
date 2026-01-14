class Solution:
    def findPoisonedDuration(self, timeSeries: List[int], duration: int) -> int:
        total_time = 0
    
        for i in range(len(timeSeries) - 1):
            
            gap = timeSeries[i+1] - timeSeries[i]
            
            total_time += min(duration, gap)
        
        return total_time + duration