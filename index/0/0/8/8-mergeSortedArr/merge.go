package main

import "log"

/**
思路:
这道题目肯定是要求原地排序的 不能额外耗内存
从后往前排列，从ab中选择比较大的从后往前放。排列完后 如果b数组剩则依次放到主数组中
因为b数组不剩的情况下不管a数组剩不剩都不需要继续排了，a数组本来就是有序的

*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	aIdx := m - 1
	bIdx := n - 1
	mainIdx := m + n - 1
	for aIdx >= 0 && bIdx >= 0 {
		if nums1[aIdx] < nums2[bIdx] {
			nums1[mainIdx] = nums2[bIdx]
			bIdx--
		} else {
			nums1[mainIdx] = nums1[aIdx]
			aIdx--
		}
		mainIdx--
	}
	for bIdx >= 0 {
		nums1[mainIdx] = nums2[bIdx]
		mainIdx--
		bIdx--
	}
}

func main() {
	arr1 := []int{1, 2, 3, 0, 0, 0}
	arr2 := []int{2, 5, 6}
	merge(arr1, 3, arr2, 3)
	log.Println(arr1)
}
