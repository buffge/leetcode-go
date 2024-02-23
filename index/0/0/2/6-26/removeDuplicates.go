package main

import "log"

/*
*
26.删除排序数组中的重复项
思路 设置索引curr为唯一元素数组的最后一个idx 比如数组[1,2,3,4,5,5,5]的唯一元素数组为[1,2,3,4] 最后一个索引就是3
初始 curr 为 0表示 [元素0] 此数组中
就是快慢指针 curr为慢指针 遍历数组的i为快指针
*/
func removeDuplicates(nums []int) int {
	nLen := len(nums)
	maxLen := 1
	for i := 0; i < nLen-1; {
		if nums[i] == nums[i+1] {
			for j := i; j < nLen-1; j++ {
				nums[j] = nums[j+1]
			}
			if nums[i] == nums[nLen-1] {
				break
			}
			continue
		}
		i++
	}
	for i := 0; i < nLen-1; i++ {
		if nums[i+1] == nums[i] {
			break
		}
		maxLen++
	}
	return maxLen
}
func removeDuplicatesV2(nums []int) int {
	curr := 0
	for _, num := range nums {
		// 从第1个值开始于当前idx比较,如果不等于就把此值加到顺序队列中,并更新当前idx
		if num != nums[curr] {
			curr++
			nums[curr] = num
		}
	}
	// 要加上第0个值
	return curr + 1
}

// 重新做一遍 加个条件限制
func removeDuplicatesV3(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	curr := 0 // 设置索引curr左边为唯一元素 如curr = 3 表示012为唯一元素
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[curr] { // 如果当前元素是唯一元素
			curr++
			nums[curr] = nums[i] // 将当前元素加入唯一元素数组中 即 将唯一元素数组扩容1 并设置最后一个元素为当前数值
		}
	}
	return curr + 1 // 数组尾元素idx+1等于数组长度
}
func main() {
	log.Println(removeDuplicates([]int{1, 2}))
	log.Println(removeDuplicatesV2([]int{1, 2}))
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	log.Println(removeDuplicatesV3(nums), nums)
}
