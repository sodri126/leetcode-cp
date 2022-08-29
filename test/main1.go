package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Case #1:", countMax([]string{"1 4", "2 3", "4 1"}))
	fmt.Println("Case #2:", countMax([]string{"2 3", "3 7", "4 1"}))
}

func countMax(upRight []string) int64 {
	matrixMap := make(map[string]int64)
	var maxValue int64
	var res int64
	for i := 0; i < len(upRight); i++ {
		rowColumn := strings.Split(upRight[i], " ")
		row, _ := strconv.Atoi(rowColumn[0])
		column, _ := strconv.Atoi(rowColumn[1])

		for i := 1; i <= row; i++ {
			for j := 1; j <= column; j++ {
				key := fmt.Sprint(i, "+", j)
				matrixMap[key]++
				if matrixMap[key] > maxValue {
					maxValue = matrixMap[key]
				}
			}
		}
	}

	for _, value := range matrixMap {
		if value == maxValue {
			res++
		}
	}
	return res
}
