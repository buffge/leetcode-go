package main

import "fmt"

/*
*
dp
*/
func canJump(nums []int) bool {
	maxIdx, end := 0, len(nums)-1
	for i := 0; i <= maxIdx; i++ {
		maxIdx = max(maxIdx, i+nums[i])
		if maxIdx >= end {
			return true
		}
	}
	return false
}
func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
