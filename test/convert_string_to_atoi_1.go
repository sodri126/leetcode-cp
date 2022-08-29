package main

import (
	"fmt"
)

func myAtoi(s string) int {
	var total int
	maxInt32 := (1 << 31) - 1
	minInt32 := (-1 << 31)

	// status
	// y = already get digit characters & still get digit characters
	// n = not already digit characters
	// t = already get digit chracters & already stopped
	isMinus := 1
	for i, status := 0, 'n'; i < len(s) && status != 't'; i++ {
		isDigit := s[i] >= '0' && s[i] <= '9'
		switch {
		case isDigit && (status == 'y' || status == 'n'):
			total = (10 * total) + int(s[i]-'0')
			status = 'y'
		case s[i] == '-' && status == 'n':
			isMinus = -1
			status = 'y'
		case s[i] == '+' && status == 'n':
			status = 'y'
		case !isDigit && status == 'y':
			status = 't'
		}
	}

	if total > maxInt32 {
		total = maxInt32
		return total
	}

	if total < minInt32 {
		total = minInt32
		return total
	}

	fmt.Println(total)
	return total * isMinus
}

func main() {
	fmt.Println(myAtoi("words and 9f87"))
}
