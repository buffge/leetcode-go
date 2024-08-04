package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("->%d ", head.Val)
		head = head.Next
	}
	fmt.Println("")
}

var removeNthFromEnd = removeNthFromEndV1

/*
双指针 P99

p1先走n步,后面p1 p2 各走一步
此时 p2与p1相距n
迭代等到p1在nil p2 就是待删除的
*/
func removeNthFromEndV1(head *ListNode, n int) *ListNode {
	padHead := &ListNode{Next: head}
	p1, p2, prev := head, head, padHead
	// p1走n步
	for i := 0; i < n-1; i++ {
		p1 = p1.Next
	}
	for {
		if p1.Next == nil { // 如果p1跑完
			prev.Next = prev.Next.Next // 将p2删除
			return padHead.Next
		}
		p1 = p1.Next
		p2 = p2.Next
		prev = prev.Next
	}
}

/*
*
优化版 将p2和prev合并为一个变量
*/
func removeNthFromEndV2(head *ListNode, n int) *ListNode {
	padHead := &ListNode{Next: head} // 构造一个首部填充用于删除头节点
	p1, prev := head, padHead
	// p1走n步
	for i := 0; i < n; i++ {
		p1 = p1.Next
	}
	// p1 p2各走1步 此时 prev.Next就是p2
	// 当p1走到结束 p2就是待删除节点
	for p1 != nil {
		p1 = p1.Next
		prev = prev.Next
	}
	prev.Next = prev.Next.Next // 将p2删除
	return padHead.Next
}
func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}},
	}
	printList(removeNthFromEndV2(head, 4))
}
