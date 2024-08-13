package main

import "log"

/*
*
思路: f(n) = f(n-1) + f(n-2)
因为在 爬到第n层前 要么在第n-1层 要么在n-2层
n-1 爬 2 到第n层  = f(n-1) 就把f(n-1)所有的尾节点后面加个->2
n-2 爬 1 到第n层  = f(n-2) 就把f(n-2)所有的尾节点后面加个->1
*/
func climbStairsV1(n int) int {
	dp := [50]int{1, 2}
	if n <= 2 {
		return dp[n-1]
	}
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
func climbStairsV2(n int) int {
	x, y, z := 0, 0, 1
	for i := 1; i <= n; i++ {
		x, y = y, z
		z = x + y
	}
	return z
}

/*
*
n(2) = 2
n(3) = 2 + 1
爬 n = (n-1) + 1 | (n-2) + 2
*/
func climbStairsS1(n int) int {
	x, y, z := 0, 0, 1 // n-2 ,n-1,n
	for i := 1; i <= n; i++ {
		x, y = y, z // n-2,n-1 = n-1,n
		z = x + y   // n  = n-1 + n-2
	}
	return z
}

var climbStairs = climbStairsS1

func main() {
	log.Println(climbStairs(3))
	log.Println(climbStairs(8))
	log.Println(climbStairs(18))
	log.Println(climbStairsV2(3))
	log.Println(climbStairsV2(8))
	log.Println(climbStairsV2(18))
}
