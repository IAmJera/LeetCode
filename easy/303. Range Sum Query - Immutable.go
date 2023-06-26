package main

import "fmt"

func main() {
	sums := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(sums.sums)
	fmt.Println(sums.SumRange(0, 2))
	fmt.Println(sums.SumRange(2, 5))
	fmt.Println(sums.SumRange(0, 5))
}

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	arr := NumArray{}
	var sum = 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		arr.sums = append(arr.sums, sum)
	}
	return arr
}

func (a *NumArray) SumRange(left int, right int) int {
	if left == 0 {
		return a.sums[right]
	}
	return a.sums[right] - a.sums[left-1]
}
