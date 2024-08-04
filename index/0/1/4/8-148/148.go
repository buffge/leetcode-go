package main

import (
	"github.com/buffge/leetcode-go/utils"
)

type (
	ListNode = utils.ListNode
)

var sortList = sortListV2

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
自顶向下归并
空间复杂度 O(logN)// 递归产生的栈开销
*/
func sortListV1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	// 快慢指针找到中间节点
	slow, fast := head, head
	var prev *ListNode
	for fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
		if fast == nil {
			break
		}
	}
	if prev != nil {
		prev.Next = nil // 去除左与右关联
	}
	// 合并左右2个有序序列
	return merge(sortList(head), sortList(slow))
}

/*
*
P99
自底向上归并 空间复杂度O(1)
5 8 9 6 7
第一轮 5 8 6 9 7, 5跟8合并 9跟6合并,7跟自己合并
第二轮 5 6 8 9 7, 58跟69合并, 7跟自己合并
第三轮 5 6 7 8 9, 5689跟7合并
*/
func sortListV2(head *ListNode) *ListNode {
	var length int
	curr := head
	for curr != nil { // 计算链表长度
		curr = curr.Next
		length++
	}
	var hairCurr, // 将归并结果集插入的位置
		left, right, // 归并时的左右组
		nextRound *ListNode // 左右组的下一轮起始节点
	hair := &ListNode{Next: head} // 定义结果
	// 从1开始归并 每轮乘2
	for subLength := 1; subLength < length; subLength <<= 1 {
		hairCurr, curr = hair, hair.Next // 结果插入到最前端
		for curr != nil {                // 将所有元素加入组中 至少一个组 最多2个组
			left = curr // 左组起始节点
			for j := subLength - 1; j > 0 && curr.Next != nil; j-- {
				curr = curr.Next // 设置左组区间
			}
			right = curr.Next // 设置右组起始节点
			curr.Next = nil   // 断开左与右
			curr = right      // 从右组开始节点开始判断
			for j := subLength - 1; j > 0 && curr != nil && curr.Next != nil; j-- {
				curr = curr.Next
			}
			nextRound = nil  // 下一轮的开始节点
			if curr != nil { // 如果右组的最后一个元素存在
				nextRound = curr.Next //  设置下一轮开始节点
				curr.Next = nil       // 断开右与下一轮
			}
			hairCurr.Next = merge(left, right) // 将结果插入到合适位置
			// 将待插入位置指针移到最右边 比如本轮该插入0位置,这轮连接了2个元素 那下一轮就该从2位置插入
			for hairCurr.Next != nil {
				hairCurr = hairCurr.Next
			}
			curr = nextRound // 从下个元素继续分组合并
		}
	}
	return hair.Next
}

func main() {
	head := &ListNode{
		Val:  5,
		Next: &ListNode{Val: 8, Next: &ListNode{Val: 9, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}},
	}
	utils.PrintList(sortList(head))
}
