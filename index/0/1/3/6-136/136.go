package main

import "fmt"

/*
异或原理
a ^ 0 = a
a ^ a ^ b = b
*/
func singleNumber(nums []int) (res int) {
	for _, num := range nums {
		res ^= num
	}
	return
}
func main() {
	fmt.Println(singleNumber([]int{1, 2, 2}))
}
