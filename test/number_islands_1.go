package main

import "fmt"

func numIslands(grid [][]byte) int {
	var totalIsland int
	mTotal := len(grid)
	for i := 0; i < mTotal; i++ {
		nTotal := len(grid[i])
		islands := make([]bool, 0)
		for j := 0; j < nTotal; j++ {
			if grid[i][j] == '0' {
				continue
			}

			top, left, bottom, right := byte('0'), byte('0'), byte('0'), byte('0')
			total := 0
			islands = append(islands, false)
			totalBoolIslands := len(islands)
			// top
			if i-1 >= 0 {
				top = grid[i-1][j]
				total++
			}

			// left
			if j-1 >= 0 {
				left = grid[i][j-1]
				total++
			}

			// bottom
			if i+1 < mTotal {
				bottom = grid[i+1][j]
				total++
			}

			// right
			if j+1 < nTotal {
				right = grid[i][j+1]
				total++
			}

			calculateIsland := int(top-'0') + int(left-'0') + int(bottom-'0') + int(right-'0')
			boolCalculateIsland := calculateIsland <= 1
			checkRow := (mTotal <= 1 && total <= 2 && calculateIsland <= 2)
			checkColumn := (nTotal <= 1 && total <= 2 && calculateIsland <= 2)
			boolRelationIslands := (totalBoolIslands-2 > 0 && !islands[totalBoolIslands-2])
			if boolCalculateIsland != (checkRow || checkColumn) != boolRelationIslands {
				fmt.Println("congrats!")
				totalIsland++
				islands[totalBoolIslands-1] = true
			}

			fmt.Println("checkRow:", checkRow)
			fmt.Println("checkColumn:", checkColumn)
			fmt.Println("boolRelationIslands", boolRelationIslands)
			fmt.Println("calculateIsland:", calculateIsland, boolCalculateIsland)
			fmt.Println("bool islands:", islands)
			fmt.Println("total:", total)
			fmt.Println("mTotal:", mTotal)
			fmt.Println("nTotal:", nTotal)
			fmt.Printf("s[%d][%d] = %s\n", i, j, string(grid[i][j]))
			fmt.Println("top:", string(top), "left:", string(left), "bottom:", string(bottom), "right:", string(right))
			fmt.Printf("got island: s[%d][%d]\n", i, j)
			fmt.Println("total island:", totalIsland)
			fmt.Println()
		}
	}
	return totalIsland
}

func main() {

}
