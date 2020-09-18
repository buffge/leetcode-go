package main

import "log"

/**
思路
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
func main() {
	log.Println(removeDuplicates([]int{1, 2}))
	log.Println(removeDuplicatesV2([]int{1, 2}))
}
