package main

func Twosum(nums []int, target int) []int {
	len := len(nums)
	mymap := make(map[int]int, len)
	for i, v := range nums {
		mymap[v] = i
	}
	var result []int
	for j, value := range nums {
		index, ok := mymap[target-value]
		if ok && index != j {
			result = []int{j, index}
			break
		}
	}
	return result
}
