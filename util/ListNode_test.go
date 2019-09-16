package util

import (
	"testing"
)

func TestListNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listNode := NewListNodeFromArr(nums)
	if !SliceEqual(listNode.ToArray(), nums) {
		t.Error("array:", nums, "--", listNode.ToArray())
	}
}
