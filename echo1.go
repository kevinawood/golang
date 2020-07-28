package main

import (
	"fmt"
	"os"
)

func main() {
	// declare 2 string variables s and sep
	var s, sep string
	// i is equal to 1 and while its less than the length of the arguments
	for i := 1; i < len(os.Args); i++ {
		// we add the space and command line arguments to the variable s
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// for initialization; condition; post
