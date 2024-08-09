package main

import (
	"bytes"
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
*
res 结果集
usedX: 那些行被放过皇后了
usedY: 第i行的皇后放在了userY[i]列上
*/
func dfs(res *[][]string, usedX []bool, usedY []int, n, row int) {
	if row == n { // 如果n行都放了皇后 说明右结果
		ans := make([]string, n)
		for x := range usedX { // 行
			arr := bytes.Repeat([]byte{'.'}, n)
			arr[usedY[x]] = 'Q' // 设置列
			ans[x] = string(arr)
		}
		*res = append(*res, ans)
		return
	}

	for y := 0; y < n; y++ { // 对第row行 尝试放在y列
		canSet := true
		for x := 0; x < row; x++ { // 判断前面所有的皇后位置 是否与当前位置冲突 row后面行还没有放过不用管
			// 如果y列冲突或者 斜轴冲突 无法防止
			if y == usedY[x] || abs(x-row) == abs(y-usedY[x]) {
				canSet = false
				break
			}
		}
		if !canSet {
			continue
		}
		usedX[row] = true
		usedY = append(usedY, y)
		dfs(res, usedX, usedY, n, row+1) // 可以放置的时候到下一行继续放
		usedX[row] = false               // 回溯
		usedY = usedY[:len(usedY)-1]     // 回溯
	}
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	dfs(&res, make([]bool, n), make([]int, 0, n), n, 0)
	return res
}
func main() {
	fmt.Println(solveNQueens(4))
}
