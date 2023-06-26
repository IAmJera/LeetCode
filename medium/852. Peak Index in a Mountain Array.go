package main

import "fmt"

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))     // 1
	fmt.Println(peakIndexInMountainArray([]int{0, 2, 1, 0}))  // 1
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2})) // 1
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] < arr[middle+1] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return left
}
