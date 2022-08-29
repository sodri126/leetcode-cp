package main

import (
	"fmt"
	"sort"
)

func MovingMedian(arr []int) []int {
	var resData []int
	windowSize := arr[0]
	for i := 1; i < len(arr); i++ {
		windowArr := make([]int, 0)
		sz := 1
		if i >= windowSize {
			sz = i - windowSize + 1
		}

		for j := i; j >= sz; j-- {
			windowArr = append(windowArr, arr[j])
		}

		sort.Ints(windowArr)

		res := 0
		totalWindowArr := len(windowArr)
		if totalWindowArr%2 == 0 {
			medianIndex := totalWindowArr / 2
			res = (windowArr[medianIndex] + windowArr[medianIndex-1]) / 2
		} else {
			res = windowArr[totalWindowArr/2]
		}
		resData = append(resData, res)
	}
	return resData
}

func main() {
	arrData := []int{5, 2, 4, 6}
	fmt.Println(MovingMedian(arrData))

	arrData = []int{3, 1, 3, 5, 10, 6, 4, 3, 1}
	fmt.Println(MovingMedian(arrData))

	arrData = []int{3, 0, 0, -2, 0, 2, 0, -2}
	fmt.Println(MovingMedian(arrData))
}
