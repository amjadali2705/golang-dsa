func islandPerimeter(grid [][]int) int {
    rows := len(grid)
	if rows == 0 {
		return 0
	}

	cols := len(grid[0])
	if cols == 0 {
		return 0
	}

	perimeter := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 1 {
				// Up
				if r == 0 || grid[r-1][c] == 0 {
					perimeter++
				}
				// Down
				if r == rows-1 || grid[r+1][c] == 0 {
					perimeter++
				}
				// Left
				if c == 0 || grid[r][c-1] == 0 {
					perimeter++
				}
				// Right
				if c == cols-1 || grid[r][c+1] == 0 {
					perimeter++
				}
			}
		}
	}
	return perimeter
}