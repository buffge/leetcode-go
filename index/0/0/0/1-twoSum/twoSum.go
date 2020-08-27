package leetcode

/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

链接：https://leetcode-cn.com/problems/two-sum
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	arrLen := len(nums)
	if arrLen < 2 {
		panic("数组长度必须大于1")
	}
	for i := 0; i < arrLen-1; i++ {
		if k, ok := m[target-nums[i]]; ok && k != i {
			return []int{i, k}
		}
	}
	return []int{}
}
