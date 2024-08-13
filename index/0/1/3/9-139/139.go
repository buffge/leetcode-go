package main

import "fmt"

/*
*
双指针dp
对于一个点 从前往后 依次取字符串判断是否存在字典中 并且字符串的开头也要是
*/
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordSet[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	n := len(s)
	for i := 1; i <= n; i++ { // 遍历指针
		for j := 0; j < i; j++ { // 截取子字符串指针
			if dp[j] && wordSet[s[j:i]] { // 如果字符串存在并且跟前面字符串相连
				dp[i] = true
			}
		}
	}
	return dp[n]
}
func main() {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wordDict))

}
