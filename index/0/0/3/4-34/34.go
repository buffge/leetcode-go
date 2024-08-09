package main

import (
	"fmt"
)

func binarySearch(nums []int, target int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return -1
}
func searchRange(nums []int, target int) []int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			minIdx, maxIdx := mid, mid
			for right := mid - 1; ; {
				idx := binarySearch(nums[lo:right+1], target)
				if idx == -1 {
					break
				}
				minIdx, right = lo+idx, lo+idx-1
			}
			for left := mid + 1; ; {
				idx := binarySearch(nums[left:hi+1], target)
				if idx == -1 {
					break
				}
				maxIdx, left = left+idx, left+idx+1
			}
			return []int{minIdx, maxIdx}

		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return []int{-1, -1}
}

func main() {
	//nums := []int{5, 7, 7, 8, 8, 10}
	nums := []int{1, 2, 3, 3, 3, 3, 4, 5, 9}
	target := 3
	fmt.Println(searchRange(nums, target))
}
