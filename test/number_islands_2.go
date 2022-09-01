package main

func numIslands1(grid [][]byte) int {
	var totalIsland int
	mTotal := len(grid)
	for i := 0; i < mTotal; i++ {
		nTotal := len(grid[i])
		for j := 0; j < nTotal; j++ {
			if grid[i][j] == '1' {
				markedIsland(grid, mTotal, nTotal, i, j)
				totalIsland++
			}
		}
	}
	return totalIsland
}

func markedIsland(grid [][]byte, mTotal, nTotal, i, j int) {
	grid[i][j] = '0'
	// top
	if i-1 >= 0 && grid[i-1][j] == '1' {
		markedIsland(grid, mTotal, nTotal, i-1, j)
	}

	// left
	if j-1 >= 0 && grid[i][j-1] == '1' {
		markedIsland(grid, mTotal, nTotal, i, j-1)
	}

	// bottom
	if i+1 < mTotal && grid[i+1][j] == '1' {
		markedIsland(grid, mTotal, nTotal, i+1, j)
	}

	// right
	if j+1 < nTotal && grid[i][j+1] == '1' {
		markedIsland(grid, mTotal, nTotal, i, j+1)
	}
}

func main() {

}
