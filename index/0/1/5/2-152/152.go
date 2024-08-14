package main

import (
	"fmt"
)

/*
最大连续乘积序列一定是去除0后从两个边界向中间的
遇见0的话 边界重算

	三种情况
	  - [负数] -   左右更小值开头
	  - [正数] -  左右皆可
	  - [0] -   左右更大值开头
*/
func maxProduct(nums []int) int {
	n := len(nums)
	var maxVal float64 // 防止溢出
	product := 1.0
	for i := 0; i < n; i++ { // 计算从左算的最大值
		product *= float64(nums[i])
		maxVal = max(maxVal, product)
		if nums[i] == 0 {
			product = 1
		}
	}
	product = 1
	for i := n - 1; i >= 0; i-- { // 计算从右算的最大值
		product *= float64(nums[i])
		maxVal = max(maxVal, product)
		if nums[i] == 0 {
			product = 1
		}
	}
	return int(maxVal)
}
func main() {
	//nums := []int{2, 3, -2, 4}
	nums := []int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}
	fmt.Println(maxProduct(nums))
}
