package main

import "log"

/*
*
思路:

计算出所有段最大序列和
返回其中最大的那个
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, currSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 计算当前段最大序列和
		currSum += nums[i]
		// 如果当前最大和小于当前值 则设置当前最大序列和为当前值
		if nums[i] > currSum {
			currSum = nums[i]
		}
		// 用当前最大和与之前所有的最大序列和比较
		if maxSum < currSum {
			maxSum = currSum
		}
	}
	// 所有序列和计算完毕后 返回其中最大的一个
	return maxSum
}

func maxSubArrayV2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxSum, sum := 0, 0
	for i := 0; i < len(nums); i++ {
		// 如果之前的序列和a为正数 那么a+nums[i]的和 一定大于 nums[i]
		// 所有设置当前最大序列和为 a+nums[i] 否则和为nums[i]
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		// 计算出前面所有段中最大和
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}

func main() {
	log.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	log.Println(maxSubArrayV2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
