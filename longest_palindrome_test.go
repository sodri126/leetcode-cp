package leetcode_cp

import (
	"testing"
)

type TestTable struct {
	Name   string
	Input  string
	Actual []string
}

func checkContains(data []string, input string) bool {
	for _, value := range data {
		if value == input {
			return true
		}
	}
	return false
}

func longestPalindrome(s string) string {
	res := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			totalCorrect := 0
			for k := i; k < j; k++ {
				if s[k] == s[j-k+i] {
					totalCorrect++
					continue
				}
				break
			}
			if totalCorrect+i == j && len(s[i:j+1]) > len(res) {
				res = s[i : j+1]
			}
		}
		if len(res) == len(s) || len(res) >= len(s)-1-i {
			break
		}
	}

	return res
}

func TestCommonSubString(t *testing.T) {
	testTable := []TestTable{
		{
			Name:   "Case#1",
			Input:  "babad",
			Actual: []string{"bab", "aba"},
		},
		{
			Name:   "Case#2",
			Input:  "cbbd",
			Actual: []string{"bb"},
		},
		{
			Name:   "Case#3",
			Input:  "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
			Actual: []string{"ranynar"},
		},
	}
	for _, value := range testTable {
		t.Run(value.Name, func(t *testing.T) {
			if !checkContains(value.Actual, longestPalindrome(value.Input)) {
				t.Fail()
			}
		})
	}

}
