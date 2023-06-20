package main

import (
	"fmt"
)

func main() {
	fmt.Println(missingNumber([]int{0, 1, 3}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{0, 1, 2, 4}))
	fmt.Println(missingNumber([]int{1, 5, 3, 2, 4, 7, 0}))
}

func missingNumber(nums []int) int {
	return len(nums)*(len(nums)+1)/2 - sum(nums)
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
