package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
p99
快慢指针 慢指针走一步,快指针走2步
如果没有环返回false
有环时快指针必然跟慢指针相遇
*/
func hasCycle(head *ListNode) bool {
	slow, fast := head, head // 快慢指针 快走2步 慢走一步
	// 判断是否有环 一直走到快慢相遇即为有环
	// 有环时一定会相遇 没环时next.next为null
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
func main() {
	loopNode := &ListNode{}
	head := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: loopNode}}},
		//Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
	}
	loopNode.Next = head
	fmt.Println(hasCycle(head))

}
