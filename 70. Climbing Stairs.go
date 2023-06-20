package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
	fmt.Println(climbStairs(3))
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	n1, n2 := 1, 2

	for i := 3; i < n+1; i++ {
		n1, n2 = n2, n2+n1
	}
	return n2
}
