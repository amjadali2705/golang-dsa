class Solution:
    def numberOfLines(self, widths: List[int], s: str) -> List[int]:
        lines = 1
        current_width = 0
        
        for char in s:
            char_idx = ord(char) - ord('a')
            char_w = widths[char_idx]
            
            if current_width + char_w > 100:
                lines += 1
                current_width = char_w
            else:
                current_width += char_w
                
        return [lines, current_width]