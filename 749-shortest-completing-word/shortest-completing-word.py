class Solution:
    def shortestCompletingWord(self, licensePlate: str, words: List[str]) -> str:
        target_counts = Counter(c.lower() for c in licensePlate if c.isalpha())
        
        words.sort(key=len)
        
        for word in words:
            word_counts = Counter(word)
            if all(word_counts[char] >= count for char, count in target_counts.items()):
                return word