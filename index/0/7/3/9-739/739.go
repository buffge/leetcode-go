package main

import "fmt"

/*
*
构造一个递减栈 用新值和递减栈顶做比较计算差距
*/
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([][2]int, 0, len(temperatures))
	for i := 0; i < len(temperatures); {
		v := temperatures[i]
		if len(stack) == 0 || v <= stack[len(stack)-1][1] {
			stack = append(stack, [2]int{i, v})
			i++
			continue
		}
		top := stack[len(stack)-1]
		res[top[0]] = i - top[0]
		stack = stack[:len(stack)-1]
	}
	return res
}
func main() {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temperatures))

}
