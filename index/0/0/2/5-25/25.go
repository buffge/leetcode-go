package main

import (
	"github.com/buffge/leetcode-go/utils"
)

type (
	ListNode = utils.ListNode
)

var reverseKGroup = reverseKGroupV1

/*
*
依次翻转 计算总数 除以k 从不需要翻转的地方再翻转一次
split: 每组的前一个元素
begin: 每组的最后一个元素

0, 1 2 3 4 5  ,k = 3
第一步翻转 到k次时与前面相连
0, 3 2 1 4 5
翻转完毕后变成
0 3 2 1 5 4
第二步将不需要翻转的地方重新翻转
即从第n/k*k次开始翻转 反正完毕后连接
0, 3 2 1 4 5
*/
func reverseKGroupV1(head *ListNode, k int) *ListNode {
	padHead := &ListNode{Next: head}
	// split: 前一组的尾 ,begin: 后一组的尾(后一组第一个进组的元素)
	split, begin, curr := padHead, padHead.Next, padHead.Next
	var prev, next *ListNode
	n := 0
	// 翻转所有
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		n++
		// 如果翻转到k次了 用上一组尾连接当前组头
		// 并将当前组尾设置为上一组的尾 下个元素作为下一组的尾
		if n%k == 0 {
			split.Next = prev
			begin.Next = curr
			split = begin
			begin = curr
			prev = nil // 避免死循环
		}
	}
	split.Next = prev // 连接2个组
	split = padHead
	times := n / k * k
	// 移动到不需要翻转的组前面
	for i := 0; i < times; i++ {
		split = split.Next
	}
	curr = split.Next
	prev = nil
	for curr != nil { // 反正不需要翻转的组
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	split.Next = prev // 连接2个组
	return padHead.Next
}

func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}},
	}
	utils.PrintList(reverseKGroup(head, 3))
}
