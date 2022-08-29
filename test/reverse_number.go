package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	var y int
	for ; x != 0; x /= 10 {
		y = (y * 10) + (x % 10)
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

func main() {
	fmt.Println(reverse(123456789))
}
