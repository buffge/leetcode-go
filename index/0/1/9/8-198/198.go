package main

import "fmt"

/*
*
dp 法
f(n) = max(v[n]+f(n-2),f(n-1))
*/
func rob(nums []int) int {
	x, y := 0, nums[0] // f(n-2),f(n-1)
	for i := 1; i < len(nums); i++ {
		if x+nums[i] > y { // 如果
			x, y = y, x+nums[i]
		} else {
			x = y
		}
	}
	return y
}
func main() {
	nums := []int{1, 1}
	fmt.Println(rob(nums))
}
