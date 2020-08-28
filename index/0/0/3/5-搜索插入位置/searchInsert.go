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
取中间值 如果大于等于目标值 二分左侧,左侧完 返回lo
必须取到最左侧一个等于目标值的才算找到索引,所以大于等于都要向左继续二分
取中间值 如果小于目标值 二分右侧 右侧完 返回hi
如果等于 返回idx
*/
func searchInsertV2(nums []int, target int) int {
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] >= target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func main() {
	log.Println(searchInsert([]int{1, 3, 5, 6}, 58))
	log.Println(searchInsertV2([]int{1, 3, 5, 6}, 3))
}
