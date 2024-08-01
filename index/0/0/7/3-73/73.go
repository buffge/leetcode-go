package main

import "log"

var setZeroes = setZeroesV1

/*
*
将所有的0行和0列记录下 再依次置0
空间复杂度O(M+N) 非最优
*/
func setZeroesV1(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	// 记录下需要置0的行和列
	zeroLines := make([]bool, len(matrix))
	zeroColumns := make([]bool, len(matrix[0]))
	for i, lines := range matrix {
		for j, v := range lines {
			// 遇到0时记录下
			if v == 0 {
				zeroLines[i] = true
				zeroColumns[j] = true
			}
		}
	}
	for i, lines := range matrix {
		for j := range lines {
			// 判断是否需要置0
			if zeroLines[i] || zeroColumns[j] {
				lines[j] = 0
			}

		}
	}
}

/*
*
P99
优化空间复杂度为O(1)即原地记录
就是借用矩阵的2个数组 并再用2个变量标记借用的这2个数组是否需要置0
*/
func setZeroesV2(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	isColZero, isLineZero := false, false
	for i, lines := range matrix {
		for j, v := range lines {
			if v == 0 {
				matrix[i][0] = 0 // 第i行需要置0
				matrix[0][j] = 0 // 第i列需要置0
				if i == 0 {      // 如果第0行存在0
					isLineZero = true
				}
				if j == 0 { // 如果第0列存在0
					isColZero = true
				}
			}
		}
	}
	// 行置0 从第一行开始
	for i := 1; i < len(matrix); i++ {
		if matrix[i][0] == 0 { // 如果行需要置0
			for j := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}
	// 列置0 从第1列开始
	for i := 1; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			for j := range matrix {
				matrix[j][i] = 0
			}
		}
	}
	// 如果第0行需要置0
	if isLineZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	// 如果第0列需要置0
	if isColZero {
		for i := range matrix {
			matrix[i][0] = 0
		}
	}
}
func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	log.Println(matrix)
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroesV2(matrix)
	log.Println(matrix)

}
