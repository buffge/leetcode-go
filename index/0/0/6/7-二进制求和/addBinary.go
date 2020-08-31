package main

import "log"

/**
思路:
计算出最大长度,
开辟数组,长度为maxLen+1,用来存放二进制数
数组首位用来进位
逆序遍历,如果a,b 此位存在则相加
如果需要进位则加1,如果>=2 则-2 进位
*/
func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	var maxLen int
	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}
	arr := make([]byte, maxLen+1)
	carry := false
	for i := maxLen; i >= 0; i-- {
		if carry {
			carry = false
			arr[i]++
		}
		arr[i] += '0'
		if i >= maxLen-aLen+1 {
			arr[i] += (a[i-maxLen+aLen-1] - '0')
		}
		if i >= maxLen-bLen+1 {
			arr[i] += (b[i-maxLen+bLen-1] - '0')
		}
		if arr[i] >= '2' {
			carry = true
			arr[i] -= 2
		}
	}
	if arr[0] == '0' {
		arr = arr[1:]
	}
	return string(arr)
}

func main() {
	log.Println(addBinary("1010", "1011"))
	log.Println(addBinary("1011", "11"))
}
