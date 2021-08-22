package main

import "fmt"

func add() func(v int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

func main() {
	var a = add()
	fmt.Println(a(1))
	fmt.Println(a(1))
}
