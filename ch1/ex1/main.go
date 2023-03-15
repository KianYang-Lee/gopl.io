package main

import (
	"fmt"
	"os"
)

// Exercisee 1.1: Modify the echo program to also print os.Args[0], the name of the command
// that invoked it.

func main() {
	s := ""
	sep := " "
	for _, arg := range os.Args {
		s += arg + sep
	}
	fmt.Println(s)
}
