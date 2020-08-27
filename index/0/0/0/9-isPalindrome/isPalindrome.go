package main

import (
	"log"
	"strconv"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-i-1] {
			return false
		}
	}
	return true
}
func main() {
	log.Println(isPalindrome(23123))
	log.Println(isPalindrome(88799788))
}
