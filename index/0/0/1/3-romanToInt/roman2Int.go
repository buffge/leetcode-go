package main

import "log"

func romanToInt(s string) int {
	sLen := len(s)
	currIdx := 0
	res := 0
	var curr uint8
	var next uint8
	for currIdx < sLen {
		curr = s[currIdx]
		next = 0
		if currIdx+1 < sLen {
			next = s[currIdx+1]
		}
		switch curr {
		case 'I':
			switch next {
			case 'V':
				res += 4
				currIdx++
			case 'X':
				res += 9
				currIdx++
			default:
				res += 1
			}
		case 'V':
			res += 5
		case 'X':
			switch next {
			case 'L':
				res += 40
				currIdx++
			case 'C':
				res += 90
				currIdx++
			default:
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			switch next {
			case 'D':
				res += 400
				currIdx++
			case 'M':
				res += 900
				currIdx++
			default:
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		}
		currIdx++
	}
	return res
}
func main() {
	log.Println(romanToInt("MCMXCIV"))
}
