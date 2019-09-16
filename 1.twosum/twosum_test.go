package main

import (
	"testing"
)

func TestTwosum(t *testing.T) {
	nums1 := []int{1, 2, 3, 4}
	target1 := 8
	nums2 := []int{1, 2, 3, 4}
	target2 := 7
	nums3 := []int{2, 2, 3, 4}
	target3 := 4
	result1 := Twosum(nums1, target1)
	result2 := Twosum(nums2, target2)
	fmt.Println(result2, 3)
	result3 := Twosum(nums3, target3)
	if result1 != nil {
		t.Error("echo :", result1, "Excepted:nil")
	}
	if !sliceEqual(result2, []int{2, 3}) {
		t.Error("echo :", result2, "Excepted:[2 3]")
	}
	if !sliceEqual(result3, []int{0, 1}) {
		t.Error("echo :", result3, "Excepted:[0 1]")
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
