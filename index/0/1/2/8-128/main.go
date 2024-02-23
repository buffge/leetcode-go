package main

import "log"

/*
* 128. 最长连续序列
思路 遍历一遍写入hash
遍历数组 依次取+1值 比如 567(3) 234(3)到4时找5 5的序列长度3 所以2的序列长度更新为6 01(2)  01再找就是2 2的序列长度为6
0的序列长度就为2+6=8
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]int, len(nums)) // 从key开始的最长序列长度是val
	for _, num := range nums {        // 初始设置所有最长序列都为1
		m[num] = 1
	}
	for num := range m { // 遍历数组
		if m[num] == 0 { // 如果num已在序列中 过
			continue
		}
		for i := 1; i <= len(nums); i++ { // 依次取当前数值的下n个数
			if m[num+i] > 0 { // 如果存在并且未被加入过序列中
				m[num] += m[num+i] // 将当前序列长度加上下一个值的序列长度比如 34序列长度2 加上567的序列长度3 即为2+3=5
				m[num+i] = 0       // 设置当前值已被加入序列中
			} else { // 如果下一个值不存在 序列终止
				break
			}
		}
	}
	maxVal := 0
	for _, val := range m { // 取出所有序列最长值
		if val != 0 {
			maxVal = max(maxVal, val)
		}
	}
	return maxVal
}

func main() {
	/**
	流程:
	9->未找到
	1->未找到
	4->5->6->7->8->9 长度6 没有10进入下一个
	7 已在序列
	3->4(6)即 长度7
	-1->0 长度2
	0,5,8-1,6 已在序列

	*/
	nums := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}
	log.Println(longestConsecutive(nums))
}
