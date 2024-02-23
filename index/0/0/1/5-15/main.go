package main

import (
	"log"
	"sort"
)

/*
*15. 三数之和
思路：先排序 将问题转换为两数之和 遍历数组 当前元素为a 头尾指针元素为bc，当a+b+c=0时插入结果集
重点在于不重复 当a+b+c=0时 将所有相同的头尾指针过滤掉
遍历时如果相同值也要跳过
*/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)                      // 排序
	res := make([][]int, 0, len(nums)/3) // 结果
	for i := 0; i < len(nums); i++ {     // 将3数之和转换为2数之和即 nums[i]+ x1+x2 从nums[i]开始依次寻找
		if nums[i] > 0 { // 因为是升序排序 此时不会有组合等于0
			return res
		}
		if i > 0 && nums[i] == nums[i-1] { // 避免重复结果
			continue
		}
		left := i + 1          // 头指针
		right := len(nums) - 1 // 尾指针
		for left < right {     // 当头尾重合时终止
			val := nums[i] + nums[left] + nums[right] // 计算3者和
			if val < 0 {                              // 如果小于0 小值取更大值
				left++
			} else if val == 0 { // 如果等于0
				res = append(res, []int{nums[i], nums[left], nums[right]}) // 插入结果集
				for left < right && nums[left] == nums[left+1] {           // 避开所有重复结果
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 避开所有重复结果
					right--
				}
				right-- // 此2个数不会再用到 向中心收敛
				left++
			} else {
				right--
			}
		}
	}
	return res
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	log.Println(threeSum(nums))

}
