class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        order_map = {char: i for i, char in enumerate(order)}
        
        for i in range(len(words) - 1):
            word1 = words[i]
            word2 = words[i + 1]
            
            for j in range(min(len(word1), len(word2))):
                char1, char2 = word1[j], word2[j]
                
                if char1 != char2:
                    if order_map[char1] > order_map[char2]:
                        return False
                    break
            else:
                if len(word1) > len(word2):
                    return False
                    
        return True