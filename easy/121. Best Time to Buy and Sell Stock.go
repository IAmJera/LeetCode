package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	min := prices[0]
	max := 0

	for _, price := range prices {
		max = maximal(max, price-min)
		min = minimal(min, price)
	}
	return max
}

func minimal(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
