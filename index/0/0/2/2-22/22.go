package main

import "fmt"

var total = 0

func dfs(res *[]string, prefix string, n, left, right int) {
	total++
	if left == n && right == n {
		*res = append(*res, prefix)
		return
	}
	if left == right { // 平衡时只能是左
		dfs(res, prefix+"(", n, left+1, right)
	} else { // 左多时 可以是左或右
		if left < n { // 还可以放左时继续遍历
			dfs(res, prefix+"(", n, left+1, right)
		}
		dfs(res, prefix+")", n, left, right+1)
	}
}

/*
*
左右配对法
*/
func generateParenthesis(n int) []string {
	res := make([]string, 0, n)
	dfs(&res, "", n, 0, 0)
	return res
}
func main() {
	fmt.Println(generateParenthesis(3), total)
}

/**
()
*/
