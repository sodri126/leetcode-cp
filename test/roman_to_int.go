package main

import "fmt"

func romanToInt(s string) int {
	var total int

	var listRoman = map[uint8]struct {
		Index int
		Value int
	}{
		'I': {
			Index: 0,
			Value: 1,
		},
		'V': {
			Index: 1,
			Value: 5,
		},
		'X': {
			Index: 2,
			Value: 10,
		},
		'L': {
			Index: 3,
			Value: 50,
		},
		'C': {
			Index: 4,
			Value: 100,
		},
		'D': {
			Index: 5,
			Value: 500,
		},
		'M': {
			Index: 6,
			Value: 1000,
		},
	}

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && listRoman[s[i]].Index%2 == 0 &&
			listRoman[s[i]].Index < len(listRoman)-1 &&
			s[i] != s[i+1] &&
			(listRoman[s[i+1]].Value-listRoman[s[i]].Value) > 0 {
			total += listRoman[s[i+1]].Value - listRoman[s[i]].Value
			i++
		} else {
			total += listRoman[s[i]].Value
		}
	}
	return total
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
