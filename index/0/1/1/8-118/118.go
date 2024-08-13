package main

import "fmt"

/*
*
杨辉三角 arr[i][j] = arr[i-1][j-1] + arr[i-1][j]
*/
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(10))
}
