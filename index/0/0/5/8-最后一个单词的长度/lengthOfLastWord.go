package main

import "log"

/**
思路:
	逆序遍历:
如果为' ' 则判断前面是否为非空格(即length!=0) -> 返回length
!-> 下一轮
如果不为' ' length++
*/
func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if length != 0 {
				return length
			}
			continue
		}
		length++
	}
	return length
}
func main() {
	log.Println(lengthOfLastWord("Hello World"))
}
