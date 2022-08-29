package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	totalNums := len(nums1)
	if totalNums%2 == 1 {
		res = float64(nums1[totalNums/2])
	} else {
		median := totalNums / 2
		res = (float64(nums1[median]) + float64(nums1[median-1])) / 2
	}
	return res
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
