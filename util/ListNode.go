package util

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNodeFromArr(listNode []int) *ListNode {
	len := len(listNode)
	if len <= 0 {
		return nil
	}
	head := &ListNode{Val: listNode[0]}
	cur := head
	for i := 1; i < len; i++ {
		cur.Next = &ListNode{Val: listNode[i]}
		cur = cur.Next
	}
	return head
}

func (listNode *ListNode) ToArray() []int {
	var arr []int
	cur := listNode
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	return arr
}

func (listNode *ListNode) Println() {
	index := 0
	cur := listNode
	for cur != nil {
		if index == 1 {
			fmt.Println(cur.Val)
		} else {
			fmt.Println(" ", cur.Val)
		}
		index++
		cur = cur.Next
	}
}
