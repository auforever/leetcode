package main

import "fmt"

func main() {
	result := lengthOfLongestSubstring("dvdf")
	fmt.Println(result)
}

func lengthOfLongestSubstring(s string) int {
	strlen := len(s)
	var sublen, max, begin int
	strmap := make(map[byte]int)
	for i := 0; i < strlen; i++ {
		again, ok := strmap[s[i]]
		strmap[s[i]] = i
		if ok {
			for j := begin; j < again; j++ {
				delete(strmap, s[j])
			}
			begin = again + 1
			if sublen > max {
				max = sublen
			}
			sublen = i - begin + 1
		} else {
			sublen++
		}
	}
	if sublen > max {
		max = sublen
	}
	return max
}
