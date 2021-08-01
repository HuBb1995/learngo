package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "yes我爱你!"
	bytes := []byte(str) //每个中文字符3个字节
	fmt.Println(utf8.RuneCountInString(str))
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Println("size: ", size)
		fmt.Printf("%c ", ch)
		fmt.Println()
	}
	fmt.Println()

	for i, ch := range []rune(str) {
		fmt.Printf("(%d %c) ", i, ch) //所有字符都是四个字节
	}
	fmt.Println()
}
