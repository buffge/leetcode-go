package main

import "fmt"

/*
*
二进制标志法 n个数用n个0/1表示
n个数有 2^n次个排列 每次排列的1就表示使用了 0 未使用
*/
func subsetsV1(nums []int) [][]int {
	var res [][]int
	for mask := 0; mask < 1<<len(nums); mask++ {
		ans := make([]int, 0, len(nums))
		for i := range nums { // mask 有n bit
			if (mask>>i)&1 == 1 { //  依次对比mask的3bit 如果bit为1表示此值被选中
				ans = append(ans, nums[i])
			}
		}
		res = append(res, ans)
	}
	return res
}

/*
*
最优解
同一个长度的n个数不会出现第二次 所以不存在次序问题
出现过的序列 + 新值 = 答案
*/
func subsetsV2(nums []int) [][]int {
	res := make([][]int, 0, len(nums)*2)
	res = append(res, nil)
	for _, num := range nums { // 第i轮 所有需要nums[i]的都已完毕
		for i, n := 0, len(res); i < n; i++ { // 用前面出现过的值+一个未出现过的num就是答案
			pre := make([]int, len(res[i]), len(res[i])+1)
			copy(pre, res[i])
			res = append(res, append(pre, num))
		}
	}
	return res
}

func dfs(nums, prefix []int, res *[][]int, curr int) {
	if len(nums) == curr { // 如果n个都设置完了
		ans := make([]int, len(prefix))
		copy(ans, prefix)
		*res = append(*res, ans)
		return
	}
	prefix = append(prefix, nums[curr])
	dfs(nums, prefix, res, curr+1)                 // 取
	dfs(nums, prefix[:len(prefix)-1], res, curr+1) // 不取
}

/*
*
子集 理解起来就是 从序列出取出独立的n个元素
n个元素只有取 不取2个状态
递归设置第i个元素取或不取 直到n个都设置完了 加入结果集
跟二进制法一样
回溯法
*/
func subsetsV3(nums []int) [][]int {
	res := make([][]int, 0, len(nums)*2)
	dfs(nums, make([]int, 0, len(nums)), &res, 0)
	return res
}

var subsets = subsetsV3

func main() {
	nums := []int{9, 0, 3, 5, 7}
	fmt.Println(subsets(nums))
}
