package main

import "fmt"

func binarySearch(nums []int, target int) int {
	fmt.Println(6, nums)
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
func search(nums []int, target int) int {
	rightVal := nums[len(nums)-1]
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi { // 找出比右边值大的最后一个数 即原来的尾 从此处再进行前后2次二分即可
		mid := lo + (hi-lo)>>1
		if nums[mid] <= rightVal {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	idx := binarySearch(nums[0:hi+1], target)
	if idx != -1 {
		return idx
	}
	idx = binarySearch(nums[hi+1:n], target)
	if idx != -1 {
		return idx + hi + 1
	}
	return -1
}
func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{3, 1}
	target := 1
	fmt.Println(search(nums, target))
}
