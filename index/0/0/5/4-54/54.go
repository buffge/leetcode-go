package main

import (
	"fmt"
)

var spiralOrder = spiralOrderV1

/*
*
方向 右下左上 循环
每次遍历前判断此方向是否可以走
走到头进入下一个方向
*/
func spiralOrderV1(matrix [][]int) []int {
	M, N := len(matrix), len(matrix[0])
	res := make([]int, 0, M*N)
	m, n := M, N
	x, y := 0, 0
	for n > 0 {
		// 向右走
		for i := 0; i < n; i++ {
			res = append(res, matrix[x][y])
			y++ // 进入下一列
		}
		// 进入下一个方向 下
		if x == M-(M-m)/2-1 { // 如果已经是最后一行
			break
		}
		m-- // 少掉一行
		x++ // 下一行
		y--
		// 向下走
		for i := 0; i < m; i++ {
			res = append(res, matrix[x][y])
			x++
		}
		// 进入下一个方向 左
		if y == (N-n)/2 { // 如果已经是最左一列
			break
		}
		n-- // 少掉一列
		y-- // 左一列
		x--
		// 向左走
		for i := 0; i < n; i++ {
			res = append(res, matrix[x][y])
			y--
		}
		// 进入下一个方向 上
		if x == (M-m)/2+1 { // 如果已经是最上一列
			break
		}
		m-- // 少掉一行
		x-- // 上一行
		y++
		// 向上走
		for i := 0; i < m; i++ {
			res = append(res, matrix[x][y])
			x--
		}
		// 进入下一个方向 右
		if x == N-(N-n)/2-1 { // 如果已经是最右一列
			break
		}
		n-- // 少掉一行
		y++ // 右一列
		x++
	}
	return res
}

// todo 可优化代码 V1太长
func spiralOrderV2(matrix [][]int) []int {
	return nil
}
func main() {
	matrix := [][]int{{1, 11}, {2, 12}, {3, 13}, {4, 14}, {5, 15}, {6, 16}, {7, 17}, {8, 18}, {9, 19}, {10, 20}}
	fmt.Println(spiralOrder(matrix))
	//fmt.Println(spiralOrderV2(matrix))

}
