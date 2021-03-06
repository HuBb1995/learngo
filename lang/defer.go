package main

import (
	"bufio"
	"fmt"
	fib2 "learngo/lang/fib"
	"os"
)

func writeFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}(file)

	writer := bufio.NewWriter(file)
	defer writer.Flush()
	f := fib2.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
