package main

import (
	"log"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastStr := "1"
	var resStr string
	var last uint8 = 0
	var curr uint8 = 0
	count := 0
	for i := 2; i <= n; i++ {
		resStr = ""
		count = 1
		last = lastStr[0]
		for j := 1; j < len(lastStr); j++ {
			curr = lastStr[j]
			if last == curr {
				count++
				continue
			}
			resStr += strconv.Itoa(count)
			resStr += string(last)
			last = curr
			count = 1
		}
		resStr += strconv.Itoa(count)
		resStr += string(last)
		lastStr = resStr
	}
	return resStr
}

func main() {
	log.Println(countAndSay(3))
	log.Println(countAndSay(5))
	log.Println(countAndSay(7))
}
