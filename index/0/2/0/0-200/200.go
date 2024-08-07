package main

import "fmt"

func isInArea(area [][]byte, x, y int) bool {
	return x >= 0 && y >= 0 && x < len(area) && y < len(area[0])
}

func dfs(area [][]byte, x, y int) {
	if !isInArea(area, x, y) { // 超出边界
		return
	}
	if area[x][y] != '1' { // 已访问过或者不是陆地
		return
	}
	area[x][y] = '2'
	dfs(area, x-1, y) // 上
	dfs(area, x+1, y) // 下
	dfs(area, x, y-1) // 左
	dfs(area, x, y+1) // 右
}

/*
*
dfs
*/
func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				res++
			}
		}
	}
	return res
}
func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(numIslands(grid))
}
