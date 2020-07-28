package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep, i := "", "", 0
	for index, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
		i = index
	}
	fmt.Println(s, i)
}
