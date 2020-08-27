package main

import (
	"log"
	"math"
	"strconv"
)

func reverse(x int) int {
	isNeg := x < 0
	y := 0
	var err error
	xStr := strconv.Itoa(x)
	if isNeg {
		xStr = xStr[1:]
	}
	xLen := len(xStr)
	yArr := make([]byte, xLen)
	for i := 0; i < xLen; i++ {
		yArr[xLen-i-1] = xStr[i]
	}
	if y, err = strconv.Atoi(string(yArr)); err == nil {
		if isNeg {
			y = -y
		}
	}
	if y > math.MaxInt32 || y < math.MinInt32 {
		y = 0
	}
	return y
}
func reverseV2(x int) int {
	var res int
	for x != 0 {
		// 溢出
		if t := int32(res); t != (t*10)/10 {
			return 0
		}
		// 进位并把最后一位设置为当前位
		res = res*10 + x%10
		x /= 10
	}
	return res
}
func main() {
	log.Println(reverse(1534236469))
	log.Println(reverseV2(1534236469))
}
