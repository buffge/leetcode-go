package main

import "fmt"

/*
*
贪心算法 区间开头去掉,不断扩展区间结尾
*/
func partitionLabels(s string) []int {
	res := make([]int, 0, 8)
	endSet := [26]int{}
	for i, c := range s {
		endSet[c-'a'] = i
	}
	start, end := 0, 0
	for i, c := range s {
		if endSet[c-'a'] > end { // 如果当前值在区间外还有 即此区间需要扩展
			end = endSet[c-'a'] // 扩展区间
		}
		if i == end { // 如果走到区间尽头 结算
			res = append(res, end-start+1)
			start = end + 1 // 下一个位置作为开始区间
		}
	}
	return res
}
func main() {
	s := "ababcbacadefegdehijhklij"
	fmt.Println(partitionLabels(s))
}
