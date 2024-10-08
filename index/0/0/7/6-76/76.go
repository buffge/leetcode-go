package main

import (
	"fmt"
	"time"
)

/*
*
滑动窗口
思路:
定义一个滑动窗口
1. 左边固定,右边向右扩展 直到窗口内有子串,记下此子串为答案
2. 从子串中找到最小子串
右边固定,左边向右扩展
每轮扩展检查窗口内是否还有子串 更新子串答案
如果子串答案为最短则直接返回
*/
var minWindow = minWindowV2

func minWindowV1(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	windowHash := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		windowHash[t[i]]++
	}
	left := 0
	right := len(t) - 1
	for i := 0; i < right; i++ {
		if _, exist := windowHash[s[i]]; exist {
			windowHash[s[i]]--
		}
	}
	res := ""
	for right < len(s) {
		c := s[right]
		if _, exist := windowHash[c]; !exist {
			right++
			continue
		}
		windowHash[c]--
		isFind := true
		for _, count := range windowHash {
			if count > 0 {
				isFind = false
				continue
			}
		}
		if !isFind {
			right++
			continue
		}
		currRes := s[left : right+1]
		if right+1-left == len(t) {
			return currRes
		}
		if len(currRes) <= len(res) || len(res) == 0 {
			res = currRes
		}
	leftLoop:
		for left <= right {
			if _, exist := windowHash[s[left]]; !exist {
				left++
				currRes = s[left : right+1]
				continue
			}
			windowHash[s[left]]++
			left++
			if right+1-left < len(t) {
				break leftLoop
			}
			if windowHash[s[left-1]] > 0 {
				break leftLoop
			}
			currRes = s[left : right+1]
		}
		if len(currRes) <= len(res) {
			res = currRes
		}
		right++
	}
	return res

}

// 普通优化
func minWindowV2(s string, t string) string {
	// 如果查找字符串长度大于被搜索字符串 返回空
	if len(s) < len(t) {
		return ""
	}
	// 定义字典 窗口内 字符->还需要的数量
	windowHash := make(map[byte]int, len(t))
	for i := 0; i < len(t); i++ {
		windowHash[t[i]]++
	}
	left := 0
	right := len(t) - 1
	// 窗口设置为0-搜索字符串长度-1
	for i := 0; i < right; i++ {
		// 每找到一个字符 减少 还需要的数量
		if _, exist := windowHash[s[i]]; exist {
			windowHash[s[i]]--
		}
	}
	// 定义结果
	res := ""
	// 窗口向右扩展
	for right < len(s) {
		c := s[right]
		// 如果当前字符串不是被搜索字符串 继续扩展
		if _, exist := windowHash[c]; !exist {
			right++
			continue
		}
		// 将 当前字符串还需要的数量-1
		windowHash[c]--
		// 判断当前是否已找到子串
		isFind := true
		for _, count := range windowHash {
			if count > 0 {
				isFind = false
				break
			}
		}
		// 未找到继续向右扩展
		if !isFind {
			right++
			continue
		}
		// 如果已找到 记录当前字符串
		currRes := s[left : right+1]
		// 如果已经最短则返回
		if right+1-left == len(t) {
			return currRes
		}
		// 如果此结果为当前最短则更新结果
		if len(currRes) <= len(res) || len(res) == 0 {
			res = currRes
		}
		// 记录当前最短子串窗口左索引
		sIdx := left
		// 将窗口左界向右扩展
	leftLoop:
		for left <= right && right+1-left >= len(t) {
			// 如果丢弃的字符不在搜索串中 继续扩展并更新 最短子串窗口左索引
			if _, exist := windowHash[s[left]]; !exist {
				left++
				sIdx = left
				continue
			}
			// 如果丢弃的字符在搜素串中
			windowHash[s[left]]++
			left++
			// 如果丢弃的字符不够了 则停止左界向右扩展,继续用右界向右扩展
			if windowHash[s[left-1]] > 0 {
				break leftLoop
			}
			// 更新最短子串窗口左索引
			sIdx = left
		}
		// 更新最短子串
		if right+1-sIdx <= len(res) {
			res = s[sIdx : right+1]
		}
		right++
	}
	return res

}

/*
  - p99优化 O(M+N)
    将for循环优化为O(1)

计算出t中所有字符的数量 并计算字符种类数
diff表示窗口中有几种字符未涵盖t
窗口向右滑动 计算是否已涵盖t
如果涵盖 则 左界向右滑动 找到最小子串并记录
找不到后继续右界向右滑动
*/
func minWindowV3(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	left := 0
	right := len(t) - 1
	tCnt := [58]int{} // 'z'-'A'+1
	winCnt := [58]int{}
	diff := 0 // 窗口中还差几种字符
	// 计算中t中字符数量和字符种类数量
	for i := 0; i < len(t); i++ {
		idx := t[i] - 'A'
		tCnt[idx]++
		if tCnt[idx] == 1 {
			diff++
		}
	}
	// 初始设窗口为0-len(t)-1 计算窗口中字符数量 并更新窗口中还差几种字符
	for i := left; i < right; i++ {
		idx := s[i] - 'A'
		winCnt[idx]++
		if winCnt[idx] == tCnt[idx] {
			diff--
		}
	}
	minLeftIdx := -1 // 假设当前没有找到子串
	minRightIdx := -1
	for right < len(s) { // 右界向右滑动
		idx := s[right] - 'A'
		winCnt[idx]++
		// 如果+1后数量相等则窗口中差的字符种类减1
		if winCnt[idx] == tCnt[idx] {
			diff--
		}
		if diff == 0 { // 如果存在子串
			// 判断是否是当前最短子串 是则更新
			if right-left <= minRightIdx-minLeftIdx || minLeftIdx == -1 {
				minLeftIdx, minRightIdx = left, right
			}

			for left <= right { // 左界右滑
				winCnt[s[left]-'A']--
				left++
				idx = s[left-1] - 'A'
				// 如果减1之前字符数量等于t并且t中字符数量大于0 表示抛弃了一个有效字符
				if winCnt[idx]+1 == tCnt[idx] && tCnt[idx] > 0 {
					diff++
					break
				}

				if right-left <= minRightIdx-minLeftIdx { // 更新子串
					minLeftIdx, minRightIdx = left, right
				}
			}
		}
		right++
	}

	if minLeftIdx != -1 { // 如果有子串则返回
		return s[minLeftIdx : minRightIdx+1]
	}
	return ""
}
func main() {
	s := "cabb"
	t := "ab"
	begin := time.Now()
	res := minWindow(s, t)
	fmt.Println(len(res)-len(t), time.Since(begin))
	begin = time.Now()
	res = minWindowV3(s, t)
	fmt.Println(len(res)-len(t), time.Since(begin))
}
