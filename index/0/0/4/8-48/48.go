package main

import (
	"fmt"
)

var rotate = rotateV1

/*
从0,0 1,1  2,2 .... n/2,n/2 到x,n-1依次转4次
*/
func rotateV1(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ { // 对角线 即各个正方形
		for j := i; j < n-i-1; j++ { // 一圈 即每边的各个块
			// top, right, bottom, left = left, top, right, bottom
			matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] =
				matrix[n-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1]
		}
	}
}
func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}
