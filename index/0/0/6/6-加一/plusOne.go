package main

import "log"

/**
思路:
数组首部加一位,预防进位
逆循环,首次设置进位 即+1
如果进位后>=10 则-10后继续进位
最后去除首部0
*/
func plusOne(digits []int) []int {
	carry := true
	digits = append([]int{0}, digits...)
	for i := len(digits) - 1; i >= 0; i-- {
		if carry {
			digits[i]++
		}
		carry = false
		if digits[i] >= 10 {
			digits[i] -= 10
			carry = true
		}
	}
	if digits[0] == 0 {
		digits = digits[1:]
	}
	return digits
}

// 只用判断循环中的最后一位是否为9就行了
func plusOneV2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i]++
			return digits
		}
	}
	digits[0] = 1
	digits = append(digits, 0)
	return digits
}
func main() {
	log.Println(plusOne([]int{1, 2, 3}))
	log.Println(plusOne([]int{9}))
	log.Println(plusOneV2([]int{1, 2, 3}))
	log.Println(plusOneV2([]int{9}))
}
