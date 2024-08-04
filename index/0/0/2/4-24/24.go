package main

import "github.com/buffge/leetcode-go/utils"

type (
	ListNode = utils.ListNode
)

/*
遍历链表当p1 p2存在时 p1,p2交换
记录p1 p2的上级prev 并再每轮交换后更新prev
*/
func swapPairs(head *ListNode) *ListNode {
	padHead := &ListNode{Next: head}
	prev := padHead
	for prev.Next != nil && prev.Next.Next != nil {
		p1, p2 := prev.Next, prev.Next.Next
		next := p2.Next
		p2.Next = p1
		p1.Next = next
		prev.Next = p2
		prev = p1
	}
	return padHead.Next
}
func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
	}
	utils.PrintList(swapPairs(head))
}
