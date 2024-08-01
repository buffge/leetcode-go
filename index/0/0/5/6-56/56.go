package main

import (
	"fmt"
	"slices"
)

var merge = mergeV1

/*
*
思路
先排序区间,确保左区间升序 这样可以确保遍历完当前区间后 当前区间不会再与任何区间关联
遍历arr 依次找到最小左区间
*/
func mergeV1(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals))
	foundFlagSet := make(map[int]bool, len(intervals))
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := 0; i < len(intervals); i++ {
		if foundFlagSet[i] {
			continue
		}
		start, end := intervals[i][0], intervals[i][1]
		for j := i + 1; j < len(intervals); j++ {
			if foundFlagSet[j] {
				continue
			}
			item := intervals[j]
			if item[0] < start {
				if item[1] < start {
					continue
				}
				foundFlagSet[j] = true
				start = item[0]
				if item[1] > end {
					end = item[1]
				}
			} else {
				if item[0] <= end {
					foundFlagSet[j] = true
					end = max(item[1], end)
				}
			}
		}
		res = append(res, []int{start, end})
	}
	return res
}

/*
*
优化代码 因为区间数组有序 左界不需要再比较,并且i可以跳过
*/
func mergeV2(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals))
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := 0; i < len(intervals); {
		start, end := intervals[i][0], intervals[i][1]
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= end {
			end = max(end, intervals[j][1])
			j++
		}
		i = j
		res = append(res, []int{start, end})
	}
	return res
}

/*
*
优化代码 改2层循环为1层循环
先取当前最小左区间,依次与后面的比较 如果在区间内合并,
不在则将当前区间加入结果,并开始下一轮
*/
func mergeV3(intervals [][]int) [][]int {
	res := make([][]int, 0, len(intervals))
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end { // 如果不相交
			res = append(res, []int{start, end})
			start, end = intervals[i][0], intervals[i][1]
		} else { // 相交则更新end
			end = max(end, intervals[i][1])
		}
	}
	res = append(res, []int{start, end})
	return res
}
func main() {
	intervals := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println(merge(intervals))
	fmt.Println(mergeV2(intervals))
	fmt.Println(mergeV3(intervals))
}
