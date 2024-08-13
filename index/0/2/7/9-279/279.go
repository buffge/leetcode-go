package main

import (
	"fmt"
)

/*
*
f(n)= f(n-1)+1
f(n)= f(n-4)+1
f(n)= f(n-9)+1
f(n)= f(n-16)+1
...
选择其中最小的
*/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minVal := n
		for j := 1; j*j <= i; j++ {
			minVal = min(minVal, dp[i-j*j])
		}
		dp[i] = minVal + 1
	}
	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
}
