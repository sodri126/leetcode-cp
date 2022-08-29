package main

import (
	"fmt"
)

func SearchingChallenge(data string) (res string) {
	k := int(data[0]) - '0'
	uniqueChar := make(map[uint8]bool)
	chars := make([]uint8, 0)
	for i := 1; i < len(data); i++ {
		if !uniqueChar[data[i]] {
			chars = append(chars, data[i])
			uniqueChar[data[i]] = true
		}
	}

	for i := 0; i < len(chars); i++ {
		for j := 1; j < k; j++ {

		}
	}
	return
}

func main() {
	fmt.Println(SearchingChallenge("2aabbacbaa"))
}
