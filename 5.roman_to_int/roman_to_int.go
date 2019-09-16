package main

import "golang.org/x/go/src/pkg/fmt"

func main() {
	ret := romanToInt("MDCCCLXXXIV")
	fmt.Println(ret)
}

func romanToInt(s string) int {
	romanToIntMap := map[byte]int{
		'I' : 1,
		'V' : 5,
		'X' : 10,
		'L' : 50,
		'C' : 100,
		'D' : 500,
		'M' : 1000,
	}
	ret := 0
	maxIndex := len(s) - 1
	for i:=0;i<=maxIndex;i++{
		//下一个字符串存在则判断是否可以联合
		nextIndex := i + 1
		if nextIndex <=maxIndex && romanToIntMap[s[nextIndex]]>romanToIntMap[s[i]] {
			ret += romanToIntMap[s[nextIndex]] - romanToIntMap[s[i]]
			i++
			continue
		}
		ret += romanToIntMap[s[i]]
	}
	return ret
}