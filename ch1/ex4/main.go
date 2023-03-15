package main

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	duplicatedFiles := make(map[string][]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ex4: %v\n", err)
				continue // Continue with for loop
			}
			countLinesAndListDuplicatedFiles(f, counts, duplicatedFiles)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
			fmt.Printf("%s\t%s\n", line, duplicatedFiles[line])
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f) // Initialize scanner
	for input.Scan() {           // While loop with condition, input.Scan read current line until EOF
		counts[input.Text()]++
	}
}

func countLinesAndListDuplicatedFiles(f *os.File, counts map[string]int, duplicatedFiles map[string][]string) {
	input := bufio.NewScanner(f) // Initialize scanner
	for input.Scan() {           // While loop with condition, input.Scan read current line until EOF
		counts[input.Text()]++
		if val, ok := duplicatedFiles[input.Text()]; !ok {
			duplicatedFiles[input.Text()] = []string{f.Name()}
		} else if !contains(val, f.Name()) {
			duplicatedFiles[input.Text()] = append(duplicatedFiles[input.Text()], f.Name())
		}
	}
}

func contains(elems []string, target string) bool {
	for _, elem := range elems {
		if target == elem {
			return true
		}
	}
	return false
}
