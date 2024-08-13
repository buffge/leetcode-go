package main

import "fmt"

/*
*
f(n) = f(n-coins[0])+1
f(n) = f(n-coins[1])+1
f(n) = f(n-coins[2])+1
选其中最小的
*/
func coinChangeV1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		minVal := amount
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				minVal = min(minVal, dp[i-coins[j]])
			}
		}
		dp[i] = minVal + 1
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

var coinChange = coinChangeV1

func main() {
	fmt.Println(coinChange([]int{474, 83, 404, 3}, 264))

}
