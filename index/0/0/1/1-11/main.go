package main

import "log"

/*
* 11. 盛最多水的容器
思路：遍历数组用第i个依次与后面的元素组成容器计算容量并更新最大值 时间复杂度 N*N
超时 此方法不对
*/
func maxArea(arr []int) int {
	maxVal := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			maxVal = max(maxVal, min(arr[i], arr[j])*(j-i))
		}
	}
	return maxVal
}

/*
* 思路: 双指针 头尾指针
先计算头尾容量
此时 头尾中长板向中心移动一格容量一定减小 因为新板就算比长板长也是无用 因为新板超出短板的长度是无用的
此时 头尾中短板向中心移动一格容量可能变大 新板比短板长就有能容量变大
所以每轮移动短板计算容量 头尾板重合时即计算完毕

移动短板时不论下一个板长还是短 所能出现的最大容量都计算过了 如果移动长板可能会有最大容量未计算到
*/
func maxAreaV2(arr []int) int {
	if len(arr) == 2 { // 如果只有两个板直接返回
		return min(arr[0], arr[1])
	}
	left := 0             // 头
	right := len(arr) - 1 // 尾
	maxVal := 0           // 最大容量
	for left < right {    // 当头尾未重合就一直计算
		// 计算当前容量并更新最大值
		maxVal = max(maxVal, min(arr[left], arr[right])*(right-left))
		//
		if left < right { // 如果头板是短板头板向右移动
			left++
		} else { // 尾板向左移动
			right--
		}
	}
	return maxVal
}
func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	log.Println(maxAreaV2(arr))
}
