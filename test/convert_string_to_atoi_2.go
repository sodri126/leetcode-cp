package main

import "fmt"

func myAtoi1(s string) int {
	var total int
	maxInt32 := (1 << 31) - 1
	minInt32 := -1 << 31

	// b for break this means for break case
	// y for already read digit characters
	// n for doesn't read digit characters
	// l for reach the limit max/min
	isMinus := 1
	status := 'n'
	for i := 0; i < len(s) && (status != 'b' && status != 'l'); i++ {
		isDigit := s[i] >= '0' && s[i] <= '9'
		characterAllowed := s[i] == '-' || s[i] == '+' || s[i] == ' '
		switch {
		case !isDigit && !characterAllowed && (status == 'n' || status == 'y'):
			status = 'b'
		case isDigit:
			total = (10 * total) + int(s[i]-'0')
			status = 'y'
			calculate := total * isMinus
			if (calculate > maxInt32 || calculate == calculate*-1) && calculate != 0 {
				if isMinus == -1 {
					total = minInt32
					status = 'l'
					break
				}
				total = maxInt32
				status = 'l'
				break
			}

			if calculate < minInt32 {
				total = minInt32
				status = 'l'
				break
			}
		case s[i] == '-':
			if status == 'y' {
				status = 'b'
				break
			}
			status = 'y'
			isMinus = -1
		case s[i] == '+':
			if status == 'y' {
				status = 'b'
				break
			}
			isMinus = 1
			status = 'y'
		case s[i] == ' ' && status == 'y':
			status = 'b'
		}
	}

	if status == 'y' || status == 'b' {
		total *= isMinus
	}

	return total
}

func main() {
	fmt.Println(myAtoi1("  -0012a42"))
}
