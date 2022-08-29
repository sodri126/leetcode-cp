package main

import "fmt"

type Characters struct {
	Index          int
	Count          int
	IsEnteredIndex bool
}

func (c *Characters) AddIndex(idx int) {
	if !c.IsEnteredIndex {
		c.Index = idx
		c.IsEnteredIndex = true
	}
}

func lengthOfLongestSubstring(s string) int {
	res := 0
	tmp := 0
	listCharacter := make(map[byte]*Characters)
	for i := 0; i < len(s); i++ {
		if _, ok := listCharacter[s[i]]; !ok {
			listCharacter[s[i]] = &Characters{}
		}

		listCharacter[s[i]].Count++
		listCharacter[s[i]].AddIndex(i)

		if listCharacter[s[i]].Count > 1 {
			newIdx := i
			if tmp > res {
				res = tmp
			}
			if len(listCharacter) > 1 {
				i = listCharacter[s[newIdx]].Index
				listCharacter = make(map[byte]*Characters)
				tmp = 0
			} else {
				listCharacter = make(map[byte]*Characters)
				listCharacter[s[newIdx]] = &Characters{}
				listCharacter[s[newIdx]].AddIndex(newIdx)
				listCharacter[s[newIdx]].Count++
				tmp = 1
			}
			continue
		}
		tmp++
	}
	if tmp > res {
		res = tmp
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}
