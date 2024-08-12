package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
*
栈方法 遇到]就结算一次
*/
func decodeString(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case ']':
			content := ""
			for {
				char := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if char == '[' {
					break
				}
				content = string(char) + content
			}
			times := ""
			for {
				char := stack[len(stack)-1]
				if char < '0' || char > '9' {
					break
				}
				stack = stack[:len(stack)-1]
				times = string(char) + times
				if len(stack) == 0 {
					break
				}
			}
			timesVal, _ := strconv.Atoi(times)
			partContent := strings.Repeat(content, timesVal)
			stack = append(stack, partContent...)
		default:
			stack = append(stack, c)
		}
	}
	return string(stack)
}
func main() {
	fmt.Println(decodeString("3[a]2[bc]"))

}
