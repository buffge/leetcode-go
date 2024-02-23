package main

import (
	"log"
	"sort"
)

/*
* 49. 字母异位词分组
思路：新建一个hash[异位词][字符串数组]
遍历数组，获取有序词 插入hash中
遍历hash 返回
*/
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))
	for _, str := range strs { // 遍历字符串
		arr := []byte(str)
		sort.Slice(arr, func(i, j int) bool { // 获取有序词
			return arr[i] < arr[j]
		})
		sortedStr := string(arr)
		if _, exist := m[sortedStr]; !exist {
			m[sortedStr] = make([]string, 0, 10)
		}
		m[sortedStr] = append(m[sortedStr], str) // 将有序词所在索引插入hash中
	}
	res := make([][]string, 0, len(m))
	for _, strList := range m { // 遍历hash
		res = append(res, strList)
	}
	return res
}
func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	log.Println(groupAnagrams(strs))
}
