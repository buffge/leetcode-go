package main

import (
	"log"
)

/*
* 42. 接雨水
思路：第一反应用矩阵做 但是空间复杂度太大为 max(高度)*n 时间复杂度 height*n*n
leetcode 超时
*/
func trap(arr []int) int {
	res := 0
	maxHeight := 0
	for _, v := range arr {
		maxHeight = max(maxHeight, v)
	}
	width := len(arr)
	matrix := make([][]bool, width)
	for i := range matrix {
		matrix[i] = make([]bool, maxHeight)
	}
	for i := 0; i < width; i++ {
		for j := 0; j < arr[i]; j++ {
			matrix[i][maxHeight-j-1] = true
		}
	}
	for i := 0; i < maxHeight; i++ {
		for j := 1; j < width-1; j++ {
			if matrix[j][maxHeight-i-1] { // 此处是砖块
				continue
			}
			leftOk := false
			rightOk := false
			for k := j - 1; k >= 0; k-- { // 左边是否有砖块
				if matrix[k][maxHeight-i-1] {
					leftOk = true
					break
				}
			}
			for k2 := j + 1; k2 < width; k2++ { // 右边是否有砖块
				if matrix[k2][maxHeight-i-1] {
					rightOk = true
					break
				}
			}
			if leftOk && rightOk {
				res++
			}
		}
	}
	//for i := 0; i < maxHeight; i++ {
	//	for j := 0; j < width; j++ {
	//		if matrix[j][i] {
	//			fmt.Printf(" x ")
	//		} else {
	//			fmt.Printf(" _ ")
	//		}
	//	}
	//	fmt.Println()
	//}
	return res
}

/*
*
思路
按行计算  从0开始 取每行左右砖块 计算中间小于层数的数量 复杂度 maxHeight*N
此方法不够好
*/
func trapV2(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}
	res := 0       // 结果
	maxHeight := 0 // 最大高度
	for _, v := range arr {
		maxHeight = max(maxHeight, v)
	}
	head := 0     // 头 从下往上 上层头尾一定不大于下层 每轮更新头尾
	tail := n - 1 // 尾
	for i := 0; i < maxHeight; i++ {
		left := head
		right := tail
		for j := left; j <= right; j++ { // 从头向右找第一个砖块并更新下一轮的头
			if arr[j] > i {
				left = j
				head = j
				break
			}
		}
		for j := right; j >= left; j-- { // 从尾向左找第一个砖块并更新下一轮的头
			if arr[j] > i {
				right = j
				tail = j
				break
			}
		}
		for left < right { // 取砖块中间的空档
			if arr[left] <= i {
				res++
			}
			left++
		}
	}
	return res
}

/*
*
思路：
单调栈 栈中从底向顶存单调递减 栈中元素为水坑或左墙
如果栈中有2个 并且下一个大于栈顶说明存在水塘  即形成了凹形状墙
遍历数组 依次与栈中元素对比 如果当前元素比栈中元素高 则有两种可能 1:栈中只有一个元素形成不了水塘
2: 栈中有1个以上元素 形成了水塘(此时水塘也可能是0高度的水塘比如 001 113)
当形成水塘时计算水塘大小并将当前水塘弹出栈 水塘宽为左墙到右墙距离(i-left-1)
水塘高为当前元素顶部到水塘的短板即 min(左墙高度,右墙高度)-arr[top]
继续用当前元素和栈中墙做成水塘计算水塘大小 直到栈空或者没有比当前元素小的即不能形成水塘则将当前元素压入栈做左墙
时间复杂度O(n) 第二轮循环中次数跟第一轮无关
*/
func trapV3(arr []int) int {
	res := 0                          // 结果
	stack := make([]int, 0, len(arr)) // 栈 保存水坑或左墙
	for right, h := range arr {       // 遍历数组 假设当前元素为右墙
		for len(stack) > 0 && h > arr[stack[len(stack)-1]] { // 当栈不为空并且当前右墙比较高
			top := stack[len(stack)-1] // 取出当前水坑 此水坑长度可能很长 但top为 水坑最左端 存在10_2 0是水坑 1是左墙 2是右墙
			// 10_2 是由 002 变成0_2 002也是一个水坑 高度为0的水坑 弹出中间0就形成了0_2
			stack = stack[:len(stack)-1] // 将当前水坑或者是当前墙弹出栈 因为它不可能成为左墙或水坑了
			if len(stack) == 0 {         // 如果没有左墙则形成不了水坑即 01 12 13这种半水塘 结束计算水塘过程
				break
			}
			left := stack[len(stack)-1]                     // 左墙索引
			width := right - left - 1                       // 水塘宽度 为 左墙到右墙的距离 如1-5中间差了3个
			height := min(arr[left], arr[right]) - arr[top] // 水塘高度为 水坑底部(arr[top])到墙的短板距离
			res += width * height
		}
		// 如果前面都不能形成水塘将当前元素压入栈中做水坑或左墙
		stack = append(stack, right)
	}
	return res
}

/*
*
按列求 计算某列能存的水只要找到左边最高和右边最高 即左墙和右墙 高度为 水坑顶部到墙短板的距离
时间复杂度 O(n²)
*/
func trapV4(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	res := 0
	for i := 1; i < len(arr)-1; i++ { // 只需要计算第二列到倒数第二列
		left := i - 1            // 设左边最高墙为左边第一个
		for j := 0; j < i; j++ { // 找到左边最高墙
			if arr[j] > arr[left] {
				left = j
			}
		}
		right := i + 1                      // 设右边最高墙为右边第一个
		for j := i + 1; j < len(arr); j++ { // 找到右边最高墙
			if arr[j] > arr[right] {
				right = j
			}
		}
		if arr[left] <= arr[i] || arr[right] <= arr[i] { // 如果无法形成水坑返回
			continue
		}
		res += min(arr[left], arr[right]) - arr[i] // 水塘大小为 1*水坑顶部到左右墙短板距离 1在代码中省略
	}
	return res
}

/*
*
思路: 优化按列求将第二轮循环优化一下 改成dp 用hash表存着每个元素左边最大墙体
时间复杂度O(n) 空间复杂度O(n)
*/
func trapV5(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	res := 0
	dpLeft := make(map[int]int, len(arr))  // 动态规划保存每个元素左边最高墙的索引
	dpRight := make(map[int]int, len(arr)) // 动态规划保存每个元素右边最高墙的索引
	maxLeft := 0
	for i := 1; i < len(arr); i++ { // 左边最高墙动态规划
		dpLeft[i] = maxLeft        // 设置当前元素左边最高墙索引
		if arr[i] > arr[maxLeft] { // 更新当前最高索引
			maxLeft = i
		}
	}
	maxRight := len(arr) - 1
	for i := len(arr) - 1; i >= 0; i-- { // 右边最高墙动态规划
		dpRight[i] = maxRight
		if arr[i] > arr[maxRight] {
			maxRight = i
		}
	}
	for i := 1; i < len(arr)-1; i++ { // 只需要计算第二列到倒数第二列
		left := dpLeft[i]
		right := dpRight[i]
		if arr[left] <= arr[i] || arr[right] <= arr[i] { // 如果无法形成水坑返回
			continue
		}
		res += min(arr[left], arr[right]) - arr[i] // 水塘大小为 1*水坑顶部到左右墙短板距离 1在代码中省略
	}
	return res
}

/*
*
思路: 优化动态规划 改进空间复杂度 先设左边第一个为左边最高墙 右边第一个为右边最高墙
比如左右最高墙 如果左比右低 那计算左2的水坑 此时左2的左墙最高为左1 右墙大于左1 所有可以得到左2的最大水量
再更新最大左右墙
时间复杂度O(n) 空间复杂度O(1)
*/
func trapV6(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	res := 0              // 结果
	left := 1             // 当前计算的左列 初始为第二个
	right := len(arr) - 2 // 当前计算的右列 初始为倒数第二个
	maxLeft := left - 1   // 最高左墙初始为第1个
	maxRight := right + 1 // 最高右墙初始为倒数第一个
	for left <= right {   // 双指针向中心合并
		if arr[maxLeft] < arr[maxRight] { // 如果左高墙小于右高墙 计算左列水滴
			if arr[maxLeft] > arr[left] { // 如果能形成水坑
				res += arr[maxLeft] - arr[left] // 增加水量
			} else { // 更新左边最高墙
				maxLeft = left
			}
			left++ // 当前左列计算完毕 过掉
		} else {
			if arr[maxRight] > arr[right] { // 如果能形成水坑
				res += arr[maxRight] - arr[right] // 增加水量
			} else { // 更新右边最高墙
				maxRight = right
			}
			right-- // 当前右列计算完毕 过掉
		}
	}
	return res
}
func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	log.Println(trap(arr))
	log.Println(trapV2(arr))
	log.Println(trapV3(arr))
	log.Println(trapV4(arr))
	log.Println(trapV5(arr))
	log.Println(trapV6(arr))
}
