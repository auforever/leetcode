package util

import (
	"testing"
)

func TestBTree(t *testing.T) {
	nums1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	nums2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	btree1 := NewBTreeFromArray(nums1)
	btree2 := NewBTreeFromArray(nums2)
	if !SliceEqual(btree1.ToArray(), nums1) {
		t.Error("array:", nums1, "--", btree1.ToArray())
	}
	if !SliceEqual(btree2.ToArray(), nums2) {
		t.Error("array:", nums2, "--", btree2.ToArray())
	}
}
