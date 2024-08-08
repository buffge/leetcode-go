package main

import "fmt"

var numMapStr = [10]string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func dfs(digits, prefix string, res *[]string, curr int) {
	if curr == len(digits) {
		if prefix != "" {
			*res = append(*res, prefix)
		}
		return
	}
	c := digits[curr]
	num := c - '0'
	dfs(digits, prefix+string(numMapStr[num][0]), res, curr+1)
	dfs(digits, prefix+string(numMapStr[num][1]), res, curr+1)
	dfs(digits, prefix+string(numMapStr[num][2]), res, curr+1)
	if num == 7 || num == 9 {
		dfs(digits, prefix+string(numMapStr[num][3]), res, curr+1)
	}
}

func letterCombinations(digits string) []string {
	res := make([]string, 0, len(digits))
	dfs(digits, "", &res, 0)
	return res
}
func main() {
	digits := ""
	fmt.Println(letterCombinations(digits))

}
