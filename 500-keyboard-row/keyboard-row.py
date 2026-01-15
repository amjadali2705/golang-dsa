class Solution:
    def findWords(self, words: List[str]) -> List[str]:
        row1 = "qwertyuiopQWERTYUIOP"
        row2 = "asdfghjklASDFGHJKL"
        row3 = "zxcvbnmZXCVBNM"

        res = []
        for word in words:
            if all(char in row1 for char in word) or \
               all(char in row2 for char in word) or \
               all(char in row3 for char in word):
                res.append(word)

        return res
        