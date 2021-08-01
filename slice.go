package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5] //针对容量，正常运行；s1[5]针对长度，会报错
	s3 := s2[2:3]
	fmt.Println(s1) // [2 3 4 5]
	fmt.Println(s2) // [5 6] slice 可扩展, 因为容量cap表示从s1首位到整个底层数组末尾的容量
	fmt.Println(s3)
	s4 := make([]int, 0xffff)
	fmt.Println(len(s4))
}
