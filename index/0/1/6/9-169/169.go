package main

import "fmt"

/*
*
消消乐法
假设nums[i]最大,遍历后面如果 还等于nums[i],+1 不等于就-1,
如果减一后等于0 设下一个数为众数
*/
func majorityElement(nums []int) int {
	res, count, n := nums[0], 1, len(nums)
	for i := 1; i < n; i++ {
		if nums[i] == res {
			count++
		} else {
			count--
			if count == 0 {
				res = nums[i+1] // 不会溢出 题目保证了存在多数元素
			}
		}
	}
	return res
}
func main() {
	fmt.Println(majorityElement([]int{1, 2, 2, 4}))
}
