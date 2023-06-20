package main

import (
	"fmt"
)

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		num := nums[i] - 1
		if nums[i] != nums[num] {
			nums[i], nums[num] = nums[num], nums[i]
			i--
		}
	}
	var dis []int
	for i, num := range nums {
		if num != i+1 {
			dis = append(dis, i+1)
		}
	}
	return dis
}
