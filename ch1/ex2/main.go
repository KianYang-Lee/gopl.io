package main

import (
	"fmt"
	"os"
)

// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments,
// one per line.

func main() {
	for i, arg := range os.Args {
		fmt.Printf("Index : %d\t Arg: %s\n", i, arg)
	}
}
