package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		totalLoop := len(res)
		if len(strs[i]) < totalLoop {
			totalLoop = len(strs[i])
		}
		pref := make([]byte, 0)
		for j := 0; j < totalLoop; j++ {
			if res[j] != strs[i][j] {
				break
			}
			pref = append(pref, res[j])
		}
		res = string(pref)
	}
	return res
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
