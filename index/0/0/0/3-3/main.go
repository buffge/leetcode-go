package main

import "log"

/*
*
3.无重复字符的最长子串
思路：双指针滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	left := 0                           // 最长子串left索引
	charSet := make(map[byte]int, 0xff) // 判断字符是否在当前序列中出现过
	res := 0                            // 结果
	for i := 0; i < len(s); i++ {       // 遍历字符串
		if idx, exist := charSet[s[i]]; idx >= left && exist { // 如果当前字符在当前序列中出现过
			left = idx + 1 // 设置最长子串left索引
		}
		charSet[s[i]] = i        // 设置当前字符所在位置
		res = max(res, i+1-left) // 更新最长子串长度
	}
	return res
}

/*
* 重做一遍 复习 思路同V1
 */
func lengthOfLongestSubstringV2(s string) int {
	m := make(map[byte]int, len(s)) // char->idx
	maxLen := 0
	curr := 0
	for i := 0; i < len(s); i++ {
		if idx, exist := m[s[i]]; exist && idx >= curr {
			curr = idx + 1
		}
		m[s[i]] = i
		maxLen = max(maxLen, i-curr+1)
	}
	return maxLen
}
func main() {
	s := "abcabcbb"
	log.Println(s, ":", lengthOfLongestSubstring(s))
	log.Println(s, ":", lengthOfLongestSubstringV2(s))
	s = "bbbbb"
	log.Println(s, ":", lengthOfLongestSubstring(s))
	log.Println(s, ":", lengthOfLongestSubstringV2(s))
	s = "pwwkew"
	log.Println(s, ":", lengthOfLongestSubstring(s))
	log.Println(s, ":", lengthOfLongestSubstringV2(s))
	s = "abba"
	log.Println(s, ":", lengthOfLongestSubstring(s))
	log.Println(s, ":", lengthOfLongestSubstringV2(s))
}
