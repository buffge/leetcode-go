package main

import (
	"fmt"
)

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

var isPalindrome = isPalindromeV1

/*
*
双指针
将链表分为长度相同的两段,前半段翻转 再依次遍历
前半段翻转再依次对比
1 2 3 4 5
a:2, b:3
a:3, b:5
*/
func isPalindromeV1(head *ListNode) bool {
	pA, pB := head, head              // 快慢指针
	var prev *ListNode                // 用于将前半段翻转
	for pB != nil && pB.Next != nil { // 让快指针走到结尾
		pB = pB.Next.Next // 快指针走2步
		next := pA.Next   // 翻转慢指针
		pA.Next = prev
		prev = pA
		pA = next // 慢指针走一步
	}
	// 此时慢指针到中间,快指针到结尾
	// 此时慢指针为后半段的起始节点
	// prev为前半段的起始节点
	if pB != nil { // 判断慢指针是否为非 null 如果是 指针长度为单数 需要将后半段右移一格 保证前后段长度一致
		pA = pA.Next
	}
	pB = prev
	// 遍历一遍2个半段 如果值不同则不是回文
	for pA != nil {
		if pA.Val != pB.Val {
			return false
		}
		pA = pA.Next
		pB = pB.Next
	}
	return true
}
func isPalindromeV2(head *ListNode) bool {
	return false
}
func main() {
	head := &ListNode{
		Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
		//Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}},
	}
	fmt.Println(isPalindrome(head))
}
