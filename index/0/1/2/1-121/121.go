package main

import "fmt"

/*
记录前面最小值 用当前值与前面最小值比较
*/
func maxProfit(prices []int) (res int) {
	minVal := prices[0]
	for i := 0; i < len(prices); i++ {
		res = max(res, prices[i]-minVal)
		minVal = min(minVal, prices[i])
	}
	return res
}
func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
