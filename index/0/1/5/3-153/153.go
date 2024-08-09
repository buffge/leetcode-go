package main

import "fmt"

func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}
	rightVal := nums[len(nums)-1]
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi { // 找出比右边值大的最后一个数 即原来的尾
		mid := lo + (hi-lo)>>1
		if nums[mid] <= rightVal {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return nums[hi+1]
}
func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
