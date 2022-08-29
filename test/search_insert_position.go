package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	idx := 0
	if high == 0 {
		idx++
		if target <= nums[high] {
			idx--
		}
		return idx
	}

	if high == 1 {
		idx = high
		if target > nums[high] {
			idx++
		}

		if target <= nums[low] && idx > 0 {
			idx--
		}

		return idx
	}

	for low <= high {
		median := (low + high) / 2
		if nums[median] == target {
			return median
		}

		if nums[median] < target {
			low = median
		} else {
			high = median
		}

		fmt.Println("median:", median)
		fmt.Println("high:", high)
		fmt.Println("low:", low)
		fmt.Println()
		if high-low == 1 {
			idx = high
			if target > nums[high] {
				idx++
			}

			if target <= nums[low] && idx > 0 {
				idx--
			}

			return idx
		}
	}

	return idx

}

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 1))
}
