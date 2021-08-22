package main

import (
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic("error score")
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	//const filename = "abc.txt"
	//if contents, err := ioutil.ReadFile(filename); err != nil {
	//	fmt.Println(err)
	//}else {
	//	fmt.Printf("%s\n", contents)
	//}
	//fmt.Println(grade(-1))
	fmt.Println(grade(10))
	fmt.Println(grade(70))
	fmt.Println(grade(85))
	fmt.Println(grade(100))
}
