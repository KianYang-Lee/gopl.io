package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

/*
Exercis e 4.2: Write a program that prints the SHA256 hash of its standard input by default but
supports a command-line flag to print the SHA384 or SHA512 hash instead.
*/

var shaType = flag.String("sh", "256", "type of sha (256, 384, or 512) for encryption")

func main() {
	flag.Parse()
	reader := bufio.NewScanner(os.Stdin)

	reader.Scan()
	input := reader.Bytes()
	if *shaType == "256" {
		checksum := sha256.Sum256(input)
		fmt.Println(checksum)
	} else if *shaType == "384" {
		checksum := sha512.Sum512_256(input)
		fmt.Println(checksum)
	} else if *shaType == "512" {
		checksum := sha512.Sum512(input)
		fmt.Println(checksum)
	}
}
