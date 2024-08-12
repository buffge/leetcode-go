package main

import "log"

/*
*

	思路

A 读取一个字符,如果为右括号 判断当前栈顶为是否为对应左括号,是则弹出栈 否则返回 false
B 如果为左括号 压入栈
C 判断栈是否为空
*/
func isValid(s string) bool {
	sLen := len(s)
	if sLen&1 == 1 {
		return false
	}
	stack := make([]byte, 0, sLen>>1)
	for _, c := range s {
		switch c {
		case '{', '[', '(':
			stack = append(stack, byte(c))
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
func isValidS1(s string) bool {
	stack := make([]int32, 0, len(s)>>1)
	for _, c := range s {
		switch c {
		case '[', '{', '(':
			stack = append(stack, c)
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		default:

		}
	}
	return len(stack) == 0
}
func main() {
	log.Println(isValid("[[][][]"))
	log.Println(isValid("[[][][]]"))
}
