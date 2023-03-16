package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an
optional sign.
*/

func comma(s string) string {
	stringArr := strings.Split(s, ".")
	var output string
	if len(stringArr) == 1 {
		output = addCommaToFront(stringArr[0])
	} else if len(stringArr) == 2 {
		front := addCommaToFront(stringArr[0])
		back := addCommaToBack(stringArr[1])
		output = strings.Join([]string{front, back}, ".")
	}

	return output
}

func addCommaToFront(s string) string {
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

func addCommaToBack(s string) string {
	buff := &bytes.Buffer{}
	posCounter := 0
	for i, _ := range s {
		buff.WriteByte(s[i])
		posCounter++
		if posCounter == 3 {
			buff.WriteByte(',')
		}
	}

	return buff.String()
}

func main() {
	fmt.Println(comma("12345"))  // 12,345
	fmt.Println(comma("1.2345")) // 1.234,5
}
