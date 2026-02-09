class Solution:
    def reverseStr(self, s: str, k: int) -> str:
        res = list(s)

        for i in range(0, len(res), 2*k):
            res[i:i+k] = res[i:i+k][::-1]

        return "".join(res)