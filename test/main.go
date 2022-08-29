package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Case 1: ", countBetween([]int32{1, 2, 2, 3, 4}, []int32{0, 2}, []int32{2, 4}))
	fmt.Println("Case 2: ", countBetween([]int32{1, 3, 5, 6, 8}, []int32{2}, []int32{6}))
	fmt.Println("Case 3: ", countBetween([]int32{4, 8, 7}, []int32{2, 4}, []int32{8, 4}))

	fmt.Println("Case 1: ", countBetween1([]int32{1, 2, 2, 3, 4}, []int32{0, 2}, []int32{2, 4}))
	fmt.Println("Case 2: ", countBetween1([]int32{1, 3, 5, 6, 8}, []int32{2}, []int32{6}))
	fmt.Println("Case 3: ", countBetween1([]int32{4, 8, 7}, []int32{2, 4}, []int32{8, 4}))
}

func countBetween(arr []int32, low []int32, high []int32) []int32 {
	totalLow := len(low)

	maximum := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > maximum {
			maximum = arr[i]
		}
	}

	data := make([]int32, maximum+1)
	for i := 0; i < len(arr); i++ {
		data[arr[i]]++
	}

	var res []int32
	for i := 0; i < totalLow; i++ {
		var r int32
		for l, h := low[i], high[i]; l <= h; l++ {
			r += data[l]
		}
		res = append(res, r)
	}
	return res
}

type Tree struct {
	val   int32
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(value int32) {
	if t == nil {
		return
	}

	if t.val == value {
		return
	}

	if t.val > value {
		if t.left == nil {
			t.left = &Tree{val: value}
			return
		}
		t.left.Insert(value)
	}

	if t.val < value {
		if t.right == nil {
			t.right = &Tree{val: value}
		}
		t.right.Insert(value)
	}
}

func countBetween1(arr []int32, low []int32, high []int32) []int32 {
	mapArr := make(map[int32]int32)
	sort.SliceIsSorted(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	for i := 0; i < len(arr); i++ {
		mapArr[arr[i]]++
	}

	var res []int32
	for i := 0; i < len(low); i++ {
		var r int32
		for arrValue, totalData := range mapArr {
			if arrValue >= low[i] && arrValue <= high[i] {
				r += totalData
			}
		}
		res = append(res, r)
	}
	return res
}

func countBetween2(arr []int32, low []int32, high []int32) []int32 {
	mapArr := make(map[int32]int32)
	//sort.SliceIsSorted(arr, func(i, j int) bool {
	//	return arr[i] < arr[j]
	//})

	tree := &Tree{val: arr[0]}
	for i := 0; i < len(arr); i++ {
		mapArr[arr[i]]++
		tree.Insert(arr[i])
	}

	var res []int32
	for i := 0; i < len(low); i++ {
		var r int32
		//for l, h := low[i], high[i]; l <= h; l++ {
		//	r += mapArr[l]
		//}
		res = append(res, r)
	}
	return res
}
