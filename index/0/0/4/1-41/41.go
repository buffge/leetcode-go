package main

import (
	"fmt"
)

var firstMissingPositive = firstMissingPositiveV1

/*
hash记录1-maxVal 再从1遍历到maxVal+1 找到第一个不存在值
空间复杂度不满足 这里新建了一个hash 空间复杂度为O(N)
*/
func firstMissingPositiveV1(nums []int) int {
	numSet := make(map[int]struct{}, len(nums))
	maxVal := 0
	for _, num := range nums {
		if num > 0 {
			numSet[num] = struct{}{}
			maxVal = max(maxVal, num)
		}
	}
	for i := 1; i <= maxVal+1; i++ {
		if _, exist := numSet[i]; !exist {
			return i
		}
	}
	return 1
}

/*
*
P99
原地hash
将[1,N]放到[0,N)上 再遍历数组 不递增的则为结果 都递增则为N+1
*/
func firstMissingPositiveV2(nums []int) int {
	// 依次将每个值放到第i-1的位置上
	for i := 0; i < len(nums); i++ {
		// 防止替换的值还没有放好
		for nums[i] > 0 && nums[i] < len(nums) && nums[i] != nums[nums[i]-1] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	// 从0开始找 第一个不对的就是结果
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	// 都对了就返回结果+1
	return len(nums) + 1
}
func main() {
	nums := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(nums))
	fmt.Println(firstMissingPositiveV2(nums))
}
