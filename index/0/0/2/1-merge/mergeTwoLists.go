package main

import "log"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
思路
A: 定义 结果链表 res
B: 定义当前节点为结果链表
C: 当a和b不同时为nil 循环遍历
D: 取出非nil节点或者更小节点挂到curr.next
E: 设置更小节点为更小节点的next
F: 设置当前节点为当前节点的next
G: 遍历完毕返回 res.Next
*/
func mergeTwoLists(a *ListNode, b *ListNode) *ListNode {
	var res = &ListNode{}      // A
	curr := res                // B
	for a != nil || b != nil { // C
		if b != nil && (a == nil || b.Val < a.Val) {
			curr.Next = b // D
			b = b.Next    // E
		} else {
			curr.Next = a
			a = a.Next
		}
		curr = curr.Next // F
	}
	return res.Next // G
}
func main() {
	a := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	b := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	log.Printf("%+v", mergeTwoLists(a, b))
}
