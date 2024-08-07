package main

import (
	"fmt"
	"time"
)

/*
*
找出第一轮坏橘子 污染完毕后 第一轮坏橘子无效了  因为他们无法再污染到别的橘子了
当没有可供污染的坏橘子就结束
*/
func orangesRottingV2(grid [][]int) int {
	var badArr [][]int
	totalGood := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				totalGood++
			} else if grid[i][j] == 2 {
				badArr = append(badArr, []int{i, j})
			}
		}
	}
	direction := [][]int{
		{-1, 0}, {1, 0}, {0, -1}, {0, 1},
	}
	res := 0
	m, n := len(grid), len(grid[0])
	for len(badArr) > 0 {
		if totalGood == 0 {
			break
		}
		tempArr := make([][]int, 0, len(badArr))
		for i := 0; i < len(badArr); i++ {
			for _, d := range direction {
				x, y := badArr[i][0]+d[0], badArr[i][1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == 1 {
					totalGood--
					grid[x][y] = 2
					tempArr = append(tempArr, []int{x, y})
				}
			}
		}
		badArr = tempArr
		res++
	}
	if totalGood > 0 { // 如果污染完毕了还有 说明原来就没有坏橘子
		return -1
	}
	return res
}

var orangesRotting = orangesRottingV2

func main() {
	gird := [][]int{
		//{0, 2},
		{2, 1, 1}, {1, 1, 0}, {0, 1, 1},
	}
	begin := time.Now()
	fmt.Println(orangesRotting(gird), time.Since(begin))
}
