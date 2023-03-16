package main

import (
	"fmt"
)

/*
Exercise 3.12: Write a function that reports whether two strings are anagrams of each other,
that is, they contain the same letters in a different order.
*/

func isAnagram(stringA string, stringB string) bool {
	if len(stringA) != len(stringB) {
		return false
	}

	countA := [26]byte{}
	countB := [26]byte{}

	for i, _ := range stringA {
		countA[stringA[i]-byte('a')]++
		countB[stringB[i]-byte('a')]++
	}

	for i, _ := range countA {
		if countA[i] != countB[i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("a", "a"))     // true
	fmt.Println(isAnagram("ada", "ada")) // true
	fmt.Println(isAnagram("ada", "ads")) // false
}
