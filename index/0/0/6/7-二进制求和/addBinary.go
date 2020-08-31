package main

import "log"

func addBinary(a string, b string) string {
	aLen := len(a)
	bLen := len(b)
	var maxLen int
	if aLen > bLen {
		maxLen = aLen
	} else {
		maxLen = bLen
	}
	arr := make([]byte, maxLen)
	arr[0] = '0'
	carry := false
	for i := maxLen - 1; i >= 0; i-- {
		if carry {
			arr[i] += '1'
		}
		if i >= maxLen-aLen {
			arr[i] += a[i-maxLen+aLen]
		}
		if i >= maxLen-bLen {
			arr[i] += b[i-maxLen+bLen]
		}
		if arr[i] > '1' {
			carry = true
			arr[i] -= '2'
		}
	}
	if arr[0] == '0' {
		arr = arr[1:]
	}
	log.Println(arr)
	return string(arr)
}
func main() {
	log.Println(addBinary("11", "1"))
	log.Println(addBinary("1011", "11"))
}
