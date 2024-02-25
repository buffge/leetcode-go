package main

import "log"

/*
* 438. 找到字符串中所有字母异位词
思路: 滑动窗口 left作为异位词的起始索引 right作为终止索引 保证left到right之间都在hash中 如果right-left=原始异位词长度说明找到了
如果当前字符不在异位词中 将left移动到下一个字符
反思: 这种应该叫双指针 用2个指针生成字符串 并不是滑动窗口
*/
func findAnagrams(s, p string) []int {
	res := make([]int, 0, len(s))
	m := make(map[byte]int, len(p))   // 异位词map
	win := make(map[byte]int, len(p)) // 滑动窗口map
	for _, c := range p {             // 初始化异位词map
		m[byte(c)]++
	}
	left := 0
	for right := 0; right < len(s); right++ {
		if _, exist := m[s[right]]; exist { // 如果当前字符在异位词中
			win[s[right]]++ // 将滑动窗口中当前字符数量+1
		} else { // 清空滑动窗口
			win = make(map[byte]int, len(p))
			left = right + 1 // 重置滑动窗口
		}
		if win[s[right]] == m[s[right]] { // 如果当前窗口字符数量等于异位词中字符数量
			if right-left == len(p)-1 { // 如果当前滑动窗口长度是否等于异位词长度 添加到结果集中
				res = append(res, left)
				win[s[left]]-- // 将左字符数量-1
				left++         // 将滑动窗口右移一位
			} else if right-left > len(p)-1 { // 如果当前测试字符串总长度大于异位词
				win[s[left]]-- // 将左字符数量-1
				left++         // 将滑动窗口右移一位
			}
		} else if win[s[right]] > m[s[right]] { // 如果当前窗口字符数量大于异位词中字符数量
			for win[s[right]] > m[s[right]] { // 保证right之前的异位词正确 比如查找abc baa 第一轮去去掉b,第二轮去掉a
				win[s[left]]-- // 将左字符数量-1
				left++         // 将滑动窗口右移一位
			}
		}
	}
	return res
}

/*
*
思路: 滑动窗口 应是固定大小的窗口 一点一点向右边滑动 匹配则计算正确
*/
func findAnagramsV2(s, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return nil
	}
	res := make([]int, 0, sLen)
	win := [26]int{}      // 滑动窗口 当前滑动窗口的结果如 a 1个 b 1个
	m := [26]int{}        // 正确结果对照组 存放着正确结果比如 a 2个  b一个
	for i, c := range p { // 初始化对照组和滑动窗口
		m[c-'a']++
		win[s[i]-'a']++
	}
	if m == win { // 如果滑动窗口等于正确结果 添加当前索引
		res = append(res, 0)
	}
	for i := 1; i <= sLen-pLen; i++ { // 从1滑到最后一个可构成pLen个元素的地方比如 pLen =3 sLen=10 从1滑到7
		win[s[i-1]-'a']--      //窗口向右滑动一位 去掉旧左边
		win[s[i+pLen-1]-'a']++ //窗口向右滑动一位 新增新右
		if win == m {          // 如果新的窗口正确 添加当前索引
			res = append(res, i)
		}
	}
	return res
}

/*
*
滑动窗口优化版 避免每次全部比较滑动窗口和正确结果集
*/
func findAnagramsV3(s, p string) []int {
	sLen := len(s)
	pLen := len(p)
	if sLen < pLen {
		return nil
	}
	res := make([]int, 0, sLen)
	diff := [26]int{}     // 存储结果差集 如 abc 和 aac 结果差集为 -1 1 0
	for i, c := range p { // 初始化对照组和滑动窗口
		diff[c-'a']++    // 正确结果集+1
		diff[s[i]-'a']-- // 滑动窗口找到-1
	}
	diffVal := 0
	for _, v := range diff {
		if v != 0 { // 如果有一个不同
			diffVal++ // 要计算完毕不能break
		}
	}
	if diffVal == 0 { // 如果滑动窗口等于正确结果 添加当前索引
		res = append(res, 0)
	}
	for i := 1; i <= sLen-pLen; i++ { // 从1滑到最后一个可构成pLen个元素的地方比如 pLen =3 sLen=10 从1滑到7
		if diff[s[i+pLen-1]-'a'] == 1 { // 如果新右加入后使得窗口从不同变得相同
			diffVal-- // 差值减一
		} else if diff[s[i+pLen-1]-'a'] == 0 { // 如果新右加入后使得窗口从相同同变得不同
			diffVal++ // 差值加一
		}
		diff[s[i+pLen-1]-'a']--     // 新右入窗口
		if diff[s[i-1]-'a'] == -1 { // 如果旧左离开后使得窗口从不同变得相同
			diffVal-- // 差值减一
		} else if diff[s[i-1]-'a'] == 0 { // 如果旧左离开后使得窗口从相同同变得不同
			diffVal++ // 差值加一
		}
		diff[s[i-1]-'a']++ // 旧左离开窗口
		if diffVal == 0 {  // 如果新的窗口正确 添加当前索引
			res = append(res, i)
		}
	}
	return res
}

/*
*
双指针优化版
左右指针之间的元素都是不多不少的 比如 寻找abc  左右之间是bc 不会出现bb 如果出现bb就要把前一个b移除 即left右移
核心 保证当前元素在左边不会多 如果多就一直弹出左边
*/
func findAnagramsV4(s, p string) []int {
	res := make([]int, 0, len(s))
	diff := [26]int{}
	// 设定需要elem个idx 比如 需要 3个a 2个b p中每有一个a elem就+1
	// 后面每加入一个元素 对面的值就-1 只要保证新加入的值为0 即当前值是够了
	// 如 需要abbc  搜索abbbca第一个值a加入等于0 第二个值b加入等于1 第三个值b加入等于0
	// 第四个值b加入等于-1 此时就要弹出a 并给a+1 即需要窗口中需要一个a
	// 减去a后b还是-1 继续弹出左边的b 此时b等于0 继续向右滑动
	// 加入c  c=0 加入a a=0 每轮等于0时判断长度是否等于pLen
	for _, c := range p { // 初始化差集 如 abb abc 差集为 0 1 -1
		diff[c-'a']++
	}
	pLen := len(p)
	left := 0
	for right := 0; right < len(s); right++ { // 依次加入滑动窗口
		diff[s[right]-'a']--         // 将新右加入win中
		for diff[s[right]-'a'] < 0 { // 如果当前元素过多 就弹出左边 极端情况就是弹出right 比如加入一个不在p中的值就要全部弹出
			diff[s[left]-'a']++ // 弹出旧左
			left++              // win左边界右移
		}
		if right-left+1 == pLen { // 如果win宽度正确
			res = append(res, left) //加入结果集
		}
	}
	return res
}
func main() {
	s := "cbaebabacd"
	p := "abc"
	log.Println(findAnagrams(s, p))
	log.Println(findAnagramsV2(s, p))
	log.Println(findAnagramsV3(s, p))
	log.Println(findAnagramsV4(s, p))
	s = "abab"
	p = "ab"
	log.Println(findAnagrams(s, p))
	log.Println(findAnagramsV2(s, p))
	log.Println(findAnagramsV3(s, p))
	log.Println(findAnagramsV4(s, p))
}
