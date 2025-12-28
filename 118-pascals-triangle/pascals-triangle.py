class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        pascal_triangle = []

        if numRows == 0:
            return pascal_triangle

        pascal_triangle.append([1])

        for i in range(1, numRows):
            prev_row = pascal_triangle[i-1]
            curr_row = [1]

            for j in range(1, len(prev_row)):
                curr_row.append(prev_row[j] + prev_row[j-1])
            
            curr_row.append(1)
            pascal_triangle.append(curr_row)

        return pascal_triangle
        