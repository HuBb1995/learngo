package main

import "fmt"

type Num struct {
	value int
}

func (n *Num) print() {
	fmt.Println(n.value)
}

func (n *Num) Set(value int) {
	n.value = value
}

func main() {
	n := Num{value: 10}
	n.Set(11)
	n.print()
}
