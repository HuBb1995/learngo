package main

import "fmt"

func main() {
	var s []int //nil切片
	for i := 0; i < 100; i += 1 {
		s = append(s, 2*i+1)
	}
	s2 := make([]int, 10, 16)
	s3 := make([]int, 10, 32)
	fmt.Println(s2)
	fmt.Println(s3)
}
