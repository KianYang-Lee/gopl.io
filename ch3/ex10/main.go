package main

import (
	"bytes"
	"fmt"
)

/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string
concatenation.
*/

func comma(s string) string {
	buff := &bytes.Buffer{}
	posCounter := 0
	for i := len(s) - 1; i >= 0; i-- {
		buff.WriteByte(s[i])
		posCounter++
		if posCounter == 3 {
			buff.WriteByte(',')
			posCounter = 0
		}
	}

	reversed := buff.String()
	var newBuff bytes.Buffer
	for i := buff.Len() - 1; i >= 0; i-- {
		newBuff.WriteByte(reversed[i])
	}

	return newBuff.String()
}

func main() {
	fmt.Println(comma("12345")) // 12,345
}
