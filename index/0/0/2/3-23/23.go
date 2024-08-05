package main

import (
	"container/heap"
	"github.com/buffge/leetcode-go/utils"
)

type (
	ListNode = utils.ListNode
)

var mergeKLists = mergeKListsV2

type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *ListNodeHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func merge(left, right *ListNode) *ListNode {
	res := &ListNode{} // 定义结果序列 依次向结果序列中加入最值
	curr := res
	for left != nil || right != nil {
		if left != nil && (right == nil || left.Val < right.Val) {
			curr.Next = left
			left = left.Next
		} else {
			curr.Next = right
			right = right.Next
		}
		curr = curr.Next
	}
	return res.Next
}

/*
*
优先队列法
取k条链中最值出来 加入小堆
弹出小堆的最值 并将弹出值的下一个值加入到小堆
堆中最多只有k个元素
*/
func mergeKListsV1(lists []*ListNode) *ListNode {
	hair := &ListNode{}
	curr := hair
	hp := (ListNodeHeap)(make([]*ListNode, 0, len(lists)))
	heap.Init(&hp)         // 定义最小堆
	for i := range lists { // 将头元素插入堆中 堆中最多同时有k个元素
		if lists[i] != nil {
			heap.Push(&hp, lists[i])
		}
	}
	for len(hp) > 0 { // 从堆中弹出并插入结果
		head := heap.Pop(&hp).(*ListNode)
		curr.Next = head
		curr = curr.Next
		if head.Next != nil {
			heap.Push(&hp, head.Next)
		}
	}
	return hair.Next
}

/*
*
自底向上两两归并
*/
func mergeKListsV2(lists []*ListNode) *ListNode {
	length := len(lists) // 当前还剩几个序列需要合并
	if length == 0 {
		return nil
	}
	k := length
	for k > 1 { // 当合并到最后一个序列时完成
		idx := 0
		for i := 0; i < k; i += 2 { // 从0开始 两两合并
			if i == k-1 { // 如果 left 就是结尾 直接返回
				lists[idx] = lists[i]
			} else { // 两两合并插入
				lists[idx] = merge(lists[i], lists[i+1]) // 依次从头插入
			}
			idx++ // 更新下一个待插入位置
		}
		k = idx //  更新当前还剩多少个待合并序列
	}
	return lists[0]
}
func main() {
	lists := []*ListNode{
		{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
		{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
		{Val: 2, Next: &ListNode{Val: 6, Next: nil}},
	}
	utils.PrintList(mergeKLists(lists))
}
