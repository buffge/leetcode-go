package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
*
P99
第一步判断有环
第二步从头节点开始遍历 慢指针也遍历 相遇时即是环入口
原理:
慢指针在环内走一圈内必相遇
设 环外长度为a,环长度为b
当快指针与慢指针相遇时 f比s多走了n*b 因为s一定会在一圈内与f相遇 所以f必然是套圈再与s相遇
已知:
f = 2s 快指针走的步长等于2个慢指针步长
f = s + nb f步长 等于s步长+n个环长
k = a + nb // 从起始到环入口 即走了a步到环口再转n圈还是环口
即计算走了k步后在哪个节点

所以:
s = nb
k = s + a // 即慢指针再走a步到环口
*/
func detectCycle(head *ListNode) *ListNode {
	// 快 慢, 头
	slow, fast, head2 := head, head, head
	for {
		if fast == nil || fast.Next == nil { // 没环返回null
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	// 有环s继续走a步
	for head2 != slow {
		head2 = head2.Next
		slow = slow.Next
	}
	return head2
}
func main() {
	loopNode := &ListNode{}
	head := &ListNode{
		//Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: loopNode}}},
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: loopNode}}}},
	}
	loopNode.Next = head.Next
	fmt.Println(detectCycle(head))
}
