package main

import (
	"fmt"
	"slices"
)

var rotate = rotateV1

/*
*
P99
第一步 优化k %= len(nums)
目的是将[0-(len-k)) 移动到尾部 即 [0-4)
将[(len-k)-len)移动到头部 即 [4-7)
第一次翻转 (4-0]到尾部,(7-4]到头部了
第二次翻转 [0-4)
第三次翻转 [4-7)
*/
func rotateV1(nums []int, k int) {
	k %= len(nums)
	slices.Reverse(nums)
	slices.Reverse(nums[:k])
	slices.Reverse(nums[k:])
}

/*
*
暴力法
大于等于len-k的 arr[i+k] = nums[i]
小于len-k的 arr[k+i-len(nums)]=nums[i]
*/
func rotateV2(nums []int, k int) {
	k %= len(nums)
	split := len(nums) - k
	arr := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i < split {
			arr[i+k] = nums[i]
		} else {
			arr[k+i-len(nums)] = nums[i]
		}
	}
	copy(nums, arr)
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
	nums = []int{1, 2, 3, 4, 5, 6, 7}
	rotateV2(nums, k)
	fmt.Println(nums)

}
