package main

import "log"

/*
* 560. 和为 K 的子数组
思路: 穷举法 O(n²)
*/
func subarraySum(nums []int, k int) int {
	res := 0
	for i := range nums { // 遍历数组 从第一个开始加
		sum := 0
		for j := i; j < len(nums); j++ { // 从当前位置加到结束
			sum += nums[j]
			if sum == k { // 如果总和等于结果 总数+1
				res++
			}
		}
	}
	return res
}

/*
*
前缀和 hashmap方法
找到索引i->j中间的和为k
即 sum[j]-sum[i-1]=k
这里 17-10 = 7
这里 20-13 = 7
10 1  2  2  2  3
10 11 13 15 17 20
*/
func subarraySumV2(nums []int, k int) int {
	var res, sum int
	m := make(map[int]int, len(nums))
	m[0] = 1 // 第一个数字就满足时用这个表示
	for _, num := range nums {
		sum += num // 更新当前sum
		// 有几个区间满足就加几次
		res += m[sum-k] // 如果前面存在区间和等于 sum-k [}<] []区间和就是sum [}区间和就是sum-k 那么<]区间和就是k了
		m[sum]++        // 将和为sum的区间数+1
	}
	return res
}
func main() {
	nums := []int{1, 1, 1}
	k := 2
	log.Println(subarraySum(nums, k))
	log.Println(subarraySumV2(nums, k))
	nums = []int{1, 2, 3}
	k = 3
	log.Println(subarraySum(nums, k))
	log.Println(subarraySumV2(nums, k))
	nums = []int{-1, -1, 1}
	k = 0
	log.Println(subarraySum(nums, k))
	log.Println(subarraySumV2(nums, k))
}
