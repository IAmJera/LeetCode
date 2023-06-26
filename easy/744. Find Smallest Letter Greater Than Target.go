package main

import "fmt"

func main() {
	fmt.Printf("%c", nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a'))
}

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	for left <= right {
		middle := left + (right-left)/2
		if letters[middle] > target {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return letters[left]
}
