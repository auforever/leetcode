package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	l3 := addTwoNumbers(l1, l2)
	fmt.Println(l3, l3.Next, l3.Next.Next, l3.Next.Next.Next)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	l3 := result
	add := 0
	if l1.Val+l2.Val >= 10 {
		l3.Val = l1.Val + l2.Val - 10
		add = 1
	} else {
		l3.Val = l1.Val + l2.Val
	}
	for {
		if l1.Next != nil && l2.Next != nil {
			l3.Next = &ListNode{Val: add}
			l1 = l1.Next
			l2 = l2.Next
			l3 = l3.Next
			if l1.Val+l2.Val+add >= 10 {
				l3.Val = l1.Val + l2.Val + add - 10
				add = 1
			} else {
				l3.Val = l1.Val + l2.Val + add
				add = 0
			}
		} else {
			break
		}
	}
	for {
		if l1.Next != nil {
			l3.Next = &ListNode{Val: add}
			l1 = l1.Next
			l3 = l3.Next
			if l1.Val+add >= 10 {
				l3.Val = l1.Val + add - 10
				add = 1
			} else {
				l3.Val = l1.Val + add
				add = 0
			}
		} else {
			break
		}
	}
	for {
		if l2.Next != nil {
			l3.Next = &ListNode{Val: add}
			l2 = l2.Next
			l3 = l3.Next
			if l2.Val+l3.Val >= 10 {
				l3.Val = l2.Val + l3.Val - 10
				add = 1
			} else {
				l3.Val = l2.Val + l3.Val
				add = 0
			}
		} else {
			break
		}
	}
	if add == 1 {
		l3.Next = &ListNode{Val: add}
	}
	return result
}
