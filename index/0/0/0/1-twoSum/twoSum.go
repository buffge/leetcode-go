package main

import "log"

/* 1. 两数之和
 * 思路 用一个hash存val对应的idx 遍历数组找到 target-currVal的索引
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, num := range nums {
		if res, exist := m[target-num]; exist {
			return []int{res, i}
		}
		m[num] = i
	}
	return nil
}
func main() {
	log.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
