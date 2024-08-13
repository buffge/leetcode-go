package main

import (
	"fmt"
	"slices"
	"sort"
)

/*
*
dp 从前往后递推 有一个小的就+1
*/
func lengthOfLISV1(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j-- {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
	}
	return slices.Max(dp)
}

/*
*
贪心算法 二分查找
*/
func lengthOfLISV2(nums []int) int {
	var res []int // 存放目前已知的最长序列
	for _, v := range nums {
		j := sort.SearchInts(res, v) //
		if j == len(res) {           // 如果 已有序列都小于v 扩展序列
			res = append(res, v)
		} else { // 产生新序列
			res[j] = v // 放到第一个大于v的前面 比如 136,新来一个4 那就变成 134
		}
	}
	return len(res) // 长度即为所有序列中最长的
}

var lengthOfLIS = lengthOfLISV2

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))

}
