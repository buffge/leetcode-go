package main

import "log"

func removeElement(nums []int, target int) int {
	curr := 0
	for _, num := range nums {
		if target != num {
			nums[curr] = num
			curr++
		}
	}
	return curr
}

func main() {
	log.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
