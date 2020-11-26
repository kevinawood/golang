// Dup1 prints the text of each line that appears more than once
// in the standard input, preceded by its count.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

	// ToDo find out how to print error mesage if no duplicates are found

	// for key, value := range counts {
	// 	if value  1 {
	// 	fmt.Println("no duplicates")
	// }
}
