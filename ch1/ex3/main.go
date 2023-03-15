package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
Exercise 1.3: Experiment to measure the difference in running time between our potentially
inefficient versions and the one that uses strings.Join. (Section 1.6 illustrates part of the
time package, and Section 11.4 shows how to write benchmark tests for systematic performance
evaluation.)
*/

func main() {
	inefficientPrint()

	efficientPrint()
}

func inefficientPrint() {
	start := time.Now()
	s := ""
	sep := " "
	for _, arg := range os.Args {
		s += arg + sep
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.8fs \t %s\n", secs, s)
}

func efficientPrint() {
	start := time.Now()
	s := strings.Join(os.Args, " ")
	secs := time.Since(start).Seconds()
	fmt.Printf("%.8fs \t %s\n", secs, s)
}
