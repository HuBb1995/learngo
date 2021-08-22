package main

import "fmt"

func lengthOfLongestSubString(str string) int {
	//字符串很长时采用[]int记录lastOccurred,可采用pprof调优性能，
	// go test -bench . -cpuprofile cpu.out
	// go tool pprof cpu.out
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

func noCoverage() {
	fmt.Println("no coverage")
}
