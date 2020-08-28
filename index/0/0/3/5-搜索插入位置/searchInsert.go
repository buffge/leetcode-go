package main

import "log"

/**
思路:
	遍历数组,如果值大于目标值 则返回当前索引
	否则索引+1
*/
func searchInsert(nums []int, target int) int {
	var k, v int
	for k, v = range nums {
		if v >= target {
			return k
		}
		k++
	}
	return k
}

/**
二分法
取中间值 如果大于目标值 二分左侧,左侧完 返回lo
取中间值 如果小于目标值 二分右侧 右侧完 返回hi
如果等于 返回idx
*/
func searchInsertV2(nums []int, target int) int {
	nLen := len(nums)
	lo, hi, mid := 0, nLen-1, nLen>>1
	for mid >= 0 {
		log.Println(lo, hi, mid)
		if mid == lo {
			return lo
		}
		if nums[mid] > target {
			hi = mid
			mid = lo + (hi-lo)>>1
		} else {
			lo = mid
			mid = lo + (hi-lo)>>1
		}
	}
	return 0
}

func main() {
	log.Println(searchInsert([]int{1, 3, 5, 6}, 58))
	log.Println(searchInsertV2([]int{1, 3, 5, 6}, 58))
}
