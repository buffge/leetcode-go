package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
)

type (
	TreeNode = utils.TreeNode
)

/*
*
从中间节点向两边扩展
栈中保存每个待添加子节点的左右索引
*/
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	length, left, right := len(nums), 0, 0
	var root, curr *TreeNode = &TreeNode{}, nil
	stack, leftStack, rightStack := []*TreeNode{root}, []int{0}, []int{length}
	for len(stack) > 0 { // 待添加子节点的节点数量大于1
		curr, stack = stack[0], stack[1:]             // 当前节点
		left, leftStack = leftStack[0], leftStack[1:] // 边界[left,right)
		right, rightStack = rightStack[0], rightStack[1:]
		mid := left + (right-left)/2 // 中间值
		curr.Val = nums[mid]
		if left < mid { // 左边还有值
			curr.Left = &TreeNode{}
			stack, leftStack, rightStack = append(stack, curr.Left), append(leftStack, left), append(rightStack, mid)
		}
		if right > mid+1 { // 右边还有值 因为right 索引不可用 所有要right前面的一个索引大于mid
			curr.Right = &TreeNode{}
			stack, leftStack, rightStack = append(stack, curr.Right), append(leftStack, mid+1), append(rightStack, right)
		}
	}
	return root
}
func main() {
	nums := []int{-10, -3, 0, 5, 9}
	fmt.Println(sortedArrayToBST(nums))
}
