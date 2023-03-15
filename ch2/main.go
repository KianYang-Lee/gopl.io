package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/popcount"
)

func main() {
	// Only take in 1 extra arg
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	input, _ := strconv.ParseUint(os.Args[1], 10, 64) // ignore error
	fmt.Println(input)
	fmt.Println(popcount.PopCount(input))
}
