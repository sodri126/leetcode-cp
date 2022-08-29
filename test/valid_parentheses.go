package main

import "fmt"

type Stack []byte

func (s *Stack) Push(val byte) {
	*s = append(*s, val)
}

func (s *Stack) Pop() byte {
	var res byte
	ss := *s
	n := len(ss) - 1
	res = ss[n]
	ss = ss[:n]
	*s = ss
	return res
}

func (s Stack) Last() byte {
	if len(s) == 0 {
		return ' '
	}
	return s[len(s)-1]
}

func isValid(s string) bool {
	var stack Stack
	listParantheses := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for i := 0; i < len(s); i++ {
		if stack.Last() == ' ' {
			stack.Push(s[i])
			continue
		}

		if data, ok := listParantheses[stack.Last()]; data == s[i] && ok {
			stack.Pop()
		} else {
			stack.Push(s[i])
		}
	}

	if len(stack) > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
