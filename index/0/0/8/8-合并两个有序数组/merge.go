package main

import "log"

/**
思路:
1.
*/
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	aIdx := 0
	bIdx := 0
	for i := 0; i < m+n; i++ {
		a := 0
		b := 0
		if m > aIdx+1 {
			a = nums1[aIdx]
		}
		if n > bIdx+1 {
			b = nums2[bIdx]
		}
		if a == 0 {
			if b != 0 {
				tmp := nums1[i]
				nums1[i] = b
				bIdx++
				nums2[bIdx] = tmp
				continue
			}
		} else {
			if b != 0 {
				if a > b {
					tmp := nums1[i]
					nums1[i] = b
					nums2[bIdx] = tmp
				}
			}
		}
	}
	log.Println(nums1)
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
