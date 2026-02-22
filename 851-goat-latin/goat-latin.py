class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        vowels = set("aeiouAEIOU")
        res = []
        
        for i, word in enumerate(sentence.split(), 1):
            if word[0] in vowels:
                new_word = word + "ma"
            else:
                new_word = word[1:] + word[0] + "ma"
            
            new_word += ("a" * i)
            res.append(new_word)
            
        return " ".join(res)