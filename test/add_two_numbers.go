package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resNodes := &ListNode{}
	var tmp *ListNode
	excessValue := 0
	for curr := resNodes; l1 != nil || l2 != nil; curr = curr.Next {
		calVal := 0
		if l1 != nil {
			calVal += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			calVal += l2.Val
			l2 = l2.Next
		}
		calVal += excessValue
		if excessValue > 0 {
			excessValue = 0
		}

		if calVal >= 10 {
			excessValue++
			calVal -= 10
		}

		curr.Next = &ListNode{Val: calVal}
		tmp = curr.Next
	}
	if excessValue > 0 {
		tmp.Next = &ListNode{Val: 1}
	}
	return resNodes.Next
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	fmt.Println(addTwoNumbers(l1, l2))
}
