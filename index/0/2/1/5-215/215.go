package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargestV3(nums []int, k int) int {
	heap.Init((*IntHeap)(&nums))
	res := -1
	for i := 0; i < k; i++ {
		res = heap.Pop((*IntHeap)(&nums)).(int)
	}
	return res
}

/*
*
快速选择法 用快排分区法 分出idx等于n-k 就找到了答案
*/
func findKthLargestV1(nums []int, k int) int {
	target := len(nums) - k
	lo, hi := 0, len(nums)-1
	for {
		idx := partition(nums, lo, hi)
		if idx == target {
			return nums[idx]
		}
		if idx > target {
			hi = idx - 1
		} else {
			lo = idx + 1
		}
	}
}

/*
*
快速分区 跟快排算法分区是一样的
*/
func partition(nums []int, lo, hi int) int {
	slow, fast := lo, lo
	pivot := rand.Intn(hi-lo+1) + lo
	nums[pivot], nums[hi] = nums[hi], nums[pivot]
	for fast < hi {
		if nums[fast] < nums[hi] {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
	nums[slow], nums[hi] = nums[hi], nums[slow]
	return slow
}

/*
*
三路快速选择
*/
func findKthLargestV2(nums []int, k int) int {
	target := len(nums) - k
	lo, hi := 0, len(nums)-1
	for {
		pair := threePartition(nums, lo, hi)
		less, great := pair[0], pair[1]
		if less <= target && target <= great {
			return nums[target]
		}
		if less > target {
			hi = less - 1
		} else if great < target {
			lo = great + 1
		}
	}
}

func threePartition(nums []int, lo, hi int) [2]int {
	j := rand.Intn(hi-lo+1) + lo
	nums[j], nums[hi] = nums[hi], nums[j]
	pivot := nums[hi]
	less, great := lo, hi
	for i := lo; i <= great; {
		if nums[i] < pivot { // 如果当前值小于基准值 放置到less处
			nums[i], nums[less] = nums[less], nums[i]
			less++
			i++
		} else if nums[i] > pivot { // 大于 放置到great处
			nums[i], nums[great] = nums[great], nums[i]
			great--
		} else {
			i++
		}
	}
	return [2]int{less, great}
}

var findKthLargest = findKthLargestV3

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
