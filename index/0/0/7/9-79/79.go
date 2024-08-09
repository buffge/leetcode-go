package main

import "fmt"

var direction = [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func dfs(matrix [][]byte, visited [][]bool, word string, curr, x, y int) bool {
	if matrix[x][y] != word[curr] { // 如果当前点不对 回溯
		return false
	}
	if curr == len(word)-1 { // 如果当前点对了并且找到完整字符串了 返回
		return true
	}
	curr++               // 下一个需要找的字符
	visited[x][y] = true // 设置当前位置被使用过
	for _, v := range direction {
		i, j := x+v[0], y+v[1]
		if i >= 0 && i < len(matrix) && j >= 0 && j < len(matrix[0]) && // 上下左右存在并且未被访问的点
			!visited[i][j] && dfs(matrix, visited, word, curr, i, j) { // 去找下一个字符
			return true
		}
	}
	visited[x][y] = false // 回溯
	return false
}
func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			if dfs(board, visited, word, 0, i, j) { // 以每个点作为起点遍历一遍
				return true
			}
		}

	}
	return false
}
func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCB"
	fmt.Println(exist(board, word))
}
