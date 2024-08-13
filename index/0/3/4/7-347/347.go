package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

type Heap [][2]int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MinHeap [][2]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.([2]int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/*
*
堆排序法
*/
func topKFrequentV1(nums []int, k int) []int {
	vMapTimes := make(map[int]int, len(nums))
	for _, num := range nums {
		vMapTimes[num]++
	}
	arr := make(Heap, 0, len(vMapTimes))
	for num, times := range vMapTimes {
		arr = append(arr, [2]int{num, times})
	}
	heap.Init(&arr)
	res := make([]int, 0, len(arr))
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&arr).([2]int)[0])
	}
	return res
}

/*
*
堆排序优化内存版
堆中只保存 前k大的数
*/
func topKFrequentV2(nums []int, k int) []int {
	vMapTimes := make(map[int]int, len(nums))
	for _, num := range nums {
		vMapTimes[num]++
	}
	arr := make(MinHeap, 0, len(vMapTimes))
	heap.Init(&arr)
	for num, times := range vMapTimes {
		if len(arr) < k {
			heap.Push(&arr, [2]int{num, times})
		} else if times > arr[0][1] { // 堆中满k个时 如果当前出现频率大于最小值则弹出并插入 保持堆中有前k大的数
			heap.Pop(&arr)
			heap.Push(&arr, [2]int{num, times})
		}
	}
	res := make([]int, 0, len(arr))
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&arr).([2]int)[0])
	}
	return res
}

/*
*
三路快速选择法
*/
func topKFrequentV3(nums []int, k int) []int {
	vMapTimes := make(map[int]int, len(nums))
	for _, num := range nums {
		vMapTimes[num]++
	}
	arr := make(MinHeap, 0, len(vMapTimes))
	for num, times := range vMapTimes {
		arr = append(arr, [2]int{num, times})
	}
	target := len(arr) - k
	lo, hi := 0, len(arr)-1
	for {
		pair := threePartition(arr, lo, hi)
		less, great := pair[0], pair[1]
		if less <= target && target <= great {
			break
		}
		if less > target {
			hi = less - 1
		} else if great < target {
			lo = great + 1
		}
	}
	res := make([]int, 0, len(arr))
	for i := target; i < len(arr); i++ {
		res = append(res, arr[i][0])
	}
	return res
}
func threePartition(nums [][2]int, lo, hi int) [2]int {
	j := rand.Intn(hi-lo+1) + lo
	nums[j], nums[hi] = nums[hi], nums[j]
	pivot := nums[hi][1]
	less, great := lo, hi
	for i := lo; i <= great; {
		if nums[i][1] < pivot { // 如果当前值小于基准值 放置到less处
			nums[i], nums[less] = nums[less], nums[i]
			less++
			i++
		} else if nums[i][1] > pivot { // 大于 放置到great处
			nums[i], nums[great] = nums[great], nums[i]
			great--
		} else {
			i++
		}
	}
	return [2]int{less, great}
}

var topKFrequent = topKFrequentV3

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
