package main

import "golang.org/x/go/src/pkg/fmt"

func main() {
	arr := []int{1,2,3,2,3,2,1}
	result := singleNumber(arr)
	fmt.Println(result)
}

func singleNumber(nums []int) int {
	ret := 0
	for _,v := range nums {
		ret ^= v
	}
	return ret
}