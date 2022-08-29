package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var xx int
	for y := x; y != 0; y /= 10 {
		xx = (xx * 10) + (y % 10)
	}

	return xx == x
}

func main() {
	fmt.Println(isPalindrome(121))
}
