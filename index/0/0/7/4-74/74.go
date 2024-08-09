package main

import "fmt"

func searchMatrixV1(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m-1
	for left <= right {
		x := left + (right-left)>>1
		lo, hi := 0, n-1
		for lo <= hi {
			y := lo + (hi-lo)>>1
			if matrix[x][y] == target {
				return true
			}
			if matrix[x][y] > target {
				hi = y - 1
			} else {
				lo = y + 1
			}
		}
		if lo == 0 {
			right = x - 1
		} else {
			left = x + 1
		}
	}
	return false
}

/*
*
矩阵转数组再进行二分法
*/
func searchMatrixV2(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	lo, hi := 0, m*n-1
	for lo <= hi {
		mid := lo + (hi-lo)>>1
		val := matrix[mid/n][mid%n]
		if val == target {
			return true
		}
		if val > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return false
}

var searchMatrix = searchMatrixV2

func main() {
	matrix := [][]int{
		{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
		//{1, 1},
	}
	target := 7
	fmt.Println(searchMatrix(matrix, target))
}
