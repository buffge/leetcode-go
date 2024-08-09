package main

import (
	"fmt"
	"sort"
)

func dfs(nums, prefix []int, used []bool, res *[][]int, depth int) {
	if depth == len(nums) { // 如果前面遍历了所有的数 加入结果集 返回
		ans := make([]int, len(prefix)) // prefix 上一层还要用 要复制出去
		copy(ans, prefix)
		*res = append(*res, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] { // 从当前未用的元素中找一个元素插入prefix中
			continue
		}

		if used[i] || // 已使用过
			// 上一轮回溯的值跟当前值一样 表示已这个值开头的都已计算完毕
			i > 0 && !used[i-1] && nums[i-1] == nums[i] {
			continue
		}
		used[i] = true
		prefix = append(prefix, nums[i])
		dfs(nums, prefix, used, res, depth+1)
		used[i] = false
		prefix = prefix[:len(prefix)-1]
	}
}

/*
*
全排列有重复值
相同的值只加入一次 因为 1,2222 和 12,222 122,2是一样的排列
在[122]循环中 计算1开头的所有序列 第一个2设为未使用过 取
设置第1个2使用过,取第二个2 因为这个2使用过了 即 他的前一个未被使用过但是和自己一样
*/
func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0, len(nums)*2)
	dfs(nums, make([]int, 0, len(nums)), make([]bool, len(nums)), &res, 0)
	return res
}
func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))

}
