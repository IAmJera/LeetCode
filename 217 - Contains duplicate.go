package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, num := range nums {
		_, b := m[num]
		if b {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}
