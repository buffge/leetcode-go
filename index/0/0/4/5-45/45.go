package main

import (
	"fmt"
)

/*
*
求第一个可到达k的点 依次向前
*/
func jumpV1(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	idxMapMaxIdx := make(map[int]int, length)
	lastMax := 0                  // 上一轮最远到达
	for i := 0; i < length; i++ { // 找到到idx的最远索引
		maxIdx := min(i+nums[i], length-1) // 本节点最远可到达 最大达到末尾
		if maxIdx > lastMax {
			idxMapMaxIdx[maxIdx] = i
			lastMax = maxIdx
		}
	}
	res := 0
	for i := length - 1; i >= 0; {
		pre := idxMapMaxIdx[i]
		res++
		if pre == 0 {
			break
		}
		for j := pre; j < length; j++ {
			if _, exist := idxMapMaxIdx[j]; exist {
				i = j
				break
			}
		}
	}
	return res
}

/*
*
一次遍历
*/
func jumpV2(nums []int) int {
	res, begin, maxVal, k := 0, 0, 0, len(nums)-1
	for i := 0; i < k; i++ {
		maxVal = max(maxVal, i+nums[i]) // 更新可以跳到的最远距离
		if i == begin {                 // 如果经过了起点 跳跃次数加1
			res++
			begin = maxVal // 下一次的起点
		}
	}
	return res
}

var jump = jumpV2

func main() {
	//nums := []int{2, 3, 1, 1, 4}
	nums := []int{4, 1, 1, 3, 1, 1, 1}
	fmt.Println(jump(nums))
}
