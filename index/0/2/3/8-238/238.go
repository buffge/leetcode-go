package main

import (
	"fmt"
)

var productExceptSelf = productExceptSelfV1

/*
*
头尾前缀积相乘

arr: [1,2,3,4]
prefix: [1,1,2,6]
suffix: [24,12,4,1]
*/
func productExceptSelfV1(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	prefix := make([]int, length)
	suffix := make([]int, length)
	prefix[0] = 1
	suffix[length-1] = 1
	for i := 1; i < length; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := length - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}
	for i := 0; i < length; i++ {
		res[i] = prefix[i] * suffix[i]
	}
	return res
}

/*
*
优化空间按复杂度为O(1) 即节省prefix 和 suffix
用prefix代替res
arr: [1,2,3,4]
prefix: [1,1,2,6]
suffix的值不再用数组保存 而是记录 直接与prefix 相乘
*/
func productExceptSelfV2(nums []int) []int {
	length := len(nums)
	// 定义数组先作为前缀积 再与后缀积相乘 即结果
	prefix := make([]int, length)
	prefix[0] = 1
	// 计算前缀积
	for i := 1; i < length; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	// 定义后缀积
	suffixProduct := 1
	for i := length - 1; i >= 0; i-- {
		// 计算结果并更新最新后缀积
		prefix[i] = prefix[i] * suffixProduct
		suffixProduct *= nums[i]
	}
	return prefix
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
	nums = []int{1, 2, 3, 4}
	fmt.Println(productExceptSelfV2(nums))
}
