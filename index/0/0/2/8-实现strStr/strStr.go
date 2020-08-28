package main

import "log"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hLen := len(haystack)
	nLen := len(needle)
	var j int
	// 最大索引就是 hLen-nLen
	for i := 0; i < (hLen-nLen)+1; {
		// 从当前idx 对比 needle
		for j = 0; j < nLen; j++ {
			// 如果有一个字符不同就进行本轮循环
			if needle[j] != haystack[i+j] {
				i++
				break
			}
		}
		// 所有的needle都能找到便返回当前索引
		if j == nLen {
			return i
		}
	}
	return -1
}
func main() {
	// log.Println(strStr("hello", "ll"))
	log.Println(strStr("mississippi", "issip"))
}
