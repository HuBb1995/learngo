package main

import (
	"fmt"
)

func main() {
	m1 := map[string]string{
		"name": "hu",
		"sex":  "boy",
	}
	m2 := make(map[string]int)
	var m3 map[string]int

	fmt.Println(m1, m2, m3)

	for k, v := range m1 {
		fmt.Println(k, v)
	}
	if name, ok := m1["name"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println(name, ok)
	}
	fmt.Println(m1)
	delete(m1, "sex")
	fmt.Println(m1)
	str := "这里是慕课是网"
	maxLen := LongestSubString(str)
	fmt.Println(maxLen)

	fmt.Println(string([]rune(str)[3]))
}

//func LongestSubString(str string) int {
//	lastOccurred := make(map[byte]int)
//	start := 0
//	maxLength := 0
//	for i, ch := range []byte(str) {
//		lastI, ok := lastOccurred[ch]
//		if ok && lastI >= start {
//			start = lastI + 1
//		}
//		if i - start + 1 > maxLength {
//			maxLength = i - start + 1
//		}
//		lastOccurred[ch] = i
//	}
//	return maxLength
//}

func LongestSubString(str string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(str) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}
