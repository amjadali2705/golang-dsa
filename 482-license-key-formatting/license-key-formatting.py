class Solution:
    def licenseKeyFormatting(self, s: str, k: int) -> str:
        s = "".join(s.split("-")).upper()
        ss = []
        count = 0

        for i in range(len(s) - 1, -1, -1):
            ss.append(s[i])
            count += 1
            if count % k == 0:
                ss.append("-")

        if ss and ss[-1] == "-":
            ss.pop()

        return "".join(ss[::-1])