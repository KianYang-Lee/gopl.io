package main

import "fmt"

// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
func reverse(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func reverseArr(arr *[5]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	reverse(a)
	fmt.Println(a) // 5,4,3,2,1

	b := [...]int{1, 2, 3, 4, 5}
	reverseArr(&b)
	fmt.Println(b)
}
