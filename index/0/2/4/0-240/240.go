package main

import (
	"fmt"
	"log"
)

var searchMatrix = searchMatrixV1
var maxVal = 0

func search(matrix [][]int, target int, leftX, leftY, m, n int) bool {
	if leftX >= len(matrix) || leftY >= len(matrix[0]) {
		return false
	}
	x, y := m/2+leftX, (n-1)/2+leftY
	//fmt.Printf("15: leftX: %d, leftY: %d,m:%d,n:%d,x:%d,y:%d ", leftX, leftY, m, n, x, y)
	//fmt.Print("val: ", matrix[x][y], "\n")
	if matrix[x][y] == target { // 中心点
		return true
	}
	if m == 1 && n == 1 {
		return false
	}
	// 进入左上
	if matrix[x][y] > target {
		//fmt.Println("进入左上")
		leftM := x - leftX + 1
		leftN := y - leftY + 1
		if leftM == 2 && leftN == 1 {
			leftM = 1
		}
		if search(matrix, target, leftX, leftY, leftM, leftN) {
			return true
		}
	}

	if x+1 < leftX+m {
		//fmt.Println("进入下")
		// 下
		if search(matrix, target, x+1, leftY, m-x+leftX-1, y-leftY+1) {
			return true
		}
	}
	if y+1 < leftY+n {
		//fmt.Println("进入右")
		// 右
		return search(matrix, target, leftX, y+1, m, leftY+n-y-1)
	}
	return false
}

/*
*
O(lgM * lgN) 递归耗时 实际很慢
先找中心点 找到结束
未找到 如果小于 丢弃左上部分 进入下半部分搜索 再进入右半部分搜索
*/
func searchMatrixV1(matrix [][]int, target int) bool {
	return search(matrix, target, 0, 0, len(matrix), len(matrix[0]))
}

/*
*
P99 O(M+N)
优化 将矩阵看做二叉排序树 从右上角开始搜索 当没有左和右时返回
*/
func searchMatrixV2(matrix [][]int, target int) bool {
	x, y := 0, len(matrix[0])-1
	for x < len(matrix) && y >= 0 {
		if matrix[x][y] == target {
			return true
		}
		if matrix[x][y] < target {
			x++
		} else {
			y--
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//matrix := [][]int{{1, 2, 3, 4, 5}}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if !searchMatrixV2(matrix, matrix[i][j]) {
				log.Fatalln("i,j", i, j, matrix[i][j])
			}
		}
	}
	fmt.Println(matrix)

	//target := 174
	//fmt.Println(searchMatrixV1(matrix, target))
}
