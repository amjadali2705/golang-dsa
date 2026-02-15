class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        dic = {"a":".-", "b":"-...", "c":"-.-.", "d":"-..", "e":".", "f":"..-.", "g":"--.", "h":"....", "i":"..", "j":".---", "k":"-.-", "l":".-..", "m":"--", "n":"-.", "o":"---", "p":".--.", "q":"--.-", "r":".-.", "s":"...", "t":"-", "u":"..-", "v":"...-", "w":".--", "x":"-..-", "y":"-.--", "z":"--.."}
        morseCode = []

        for i in words:
            s = ""
            for j in i:
                s += dic[j]
            morseCode.append(s)

        return len(set(morseCode)) 
