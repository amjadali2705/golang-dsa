class Solution:
    def findRelativeRanks(self, score: List[int]) -> List[str]:
       
        sorted_scores = sorted(score, reverse=True)
        ranks_lookup = {}
        
        for i in range(len(sorted_scores)):
            
            if i == 0:
                ranks_lookup[sorted_scores[i]] = "Gold Medal"
            elif i == 1:
                ranks_lookup[sorted_scores[i]] = "Silver Medal"
            elif i == 2:
                ranks_lookup[sorted_scores[i]] = "Bronze Medal"
            else:
                ranks_lookup[sorted_scores[i]] = str(i + 1)
        
        res = []
        for s in score:
            res.append(ranks_lookup[s])
            
        return res