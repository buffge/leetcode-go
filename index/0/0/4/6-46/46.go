package main

import "fmt"

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
		used[i] = true
		prefix = append(prefix, nums[i])
		dfs(nums, prefix, used, res, depth+1)
		used[i] = false
		prefix = prefix[:len(prefix)-1]
	}
}

/*
回溯法 递归
*/
func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}
	res := make([][]int, 0, len(nums)*2)
	dfs(nums, make([]int, 0, len(nums)), make([]bool, len(nums)), &res, 0)
	return res
}
func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}
