package main

import (
	"container/heap"
	"fmt"
)

type Heap struct {
	arr   []int
	isMin bool
}

func (h Heap) Len() int { return len(h.arr) }

func (h Heap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Push(x any) {
	h.arr = append(h.arr, x.(int))
}

func (h *Heap) Pop() any {
	old := h.arr
	n := len(old)
	x := old[n-1]
	h.arr = old[0 : n-1]
	return x
}
func (h Heap) Less(i, j int) bool {
	if h.isMin {
		return h.arr[i] < h.arr[j]
	}
	return h.arr[i] > h.arr[j]

}

type MedianFinder struct {
	left  *Heap // 递减最大堆
	right *Heap // 递增最小堆
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &Heap{isMin: false},
		right: &Heap{isMin: true},
	}
}

func (f *MedianFinder) AddNum(num int) {
	if f.left.Len() == f.right.Len() { // 双数插入左堆
		// 如果当前值小于等于最小堆插入最大堆
		if f.left.Len() == 0 || num <= f.right.arr[0] {
			heap.Push(f.left, num)
			return
		}
		// 如果当前值大于右堆 则将右堆弹出插入到左堆 当前值插入右堆 保证右边大于左边
		heap.Push(f.left, heap.Pop(f.right))
		heap.Push(f.right, num)
	} else { // 单数插入右堆
		if num >= f.left.arr[0] {
			heap.Push(f.right, num)
			return
		}
		// 如果当前值小于左堆 将左堆弹到右堆 并将当前值插入左堆
		heap.Push(f.right, heap.Pop(f.left))
		heap.Push(f.left, num)
	}
}

func (f *MedianFinder) FindMedian() float64 {
	if f.left.Len() == f.right.Len() {
		return float64(f.left.arr[0]+f.right.arr[0]) / 2.0
	}
	return float64(f.left.arr[0])
}

func main() {
	obj := Constructor()
	obj.AddNum(6)
	obj.AddNum(10)
	obj.AddNum(2)
	obj.AddNum(6)
	obj.AddNum(5)
	obj.AddNum(0)
	obj.AddNum(6)
	obj.AddNum(3)
	obj.AddNum(1)
	obj.AddNum(0)
	obj.AddNum(0)
	// 0 0 0 1 2 3 5 6 6 6 10
	fmt.Println(obj.FindMedian())
}
