package main

import "fmt"

func dfs(nums, prefix []int, res *[][]int, sumVal, target, currIdx int) {
	if currIdx == len(nums) {
		return
	}
	if sumVal == target {
		*res = append(*res, append([]int(nil), prefix...))
		return
	}
	if sumVal < target { // 如果大于后面都不用加了 因为都是正数  小于有2种情况 +自身 继续检查遍历, 不加自身 继续跟后面的遍历
		dfs(nums, append(prefix, nums[currIdx]), res, sumVal+nums[currIdx], target, currIdx)
		dfs(nums, prefix, res, sumVal, target, currIdx+1)
	}
}

/*
*
自底向上 回溯
*/
func combinationSum(candidates []int, target int) (res [][]int) {
	dfs(candidates, nil, &res, 0, target, 0)
	return res
}

func main() {
	candidates := []int{8, 7, 4, 3}
	target := 11
	fmt.Println(combinationSum(candidates, target))
}
