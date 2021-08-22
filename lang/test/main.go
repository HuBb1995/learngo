package main

import (
	"fmt"
)

func main() {
	s := []int{1, 3, 5, 7, 9}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[0])
		s = s[1:]
	}
}
