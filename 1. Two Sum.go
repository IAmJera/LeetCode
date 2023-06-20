package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	for i, val := range nums {
		for j, value := range nums {
			if i != j && val+value == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
