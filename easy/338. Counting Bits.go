package main

import (
	"fmt"
)

func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}

func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		res[i] = res[i>>1] + i%2
	}
	return res
}
