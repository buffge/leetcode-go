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
递归法
*/
func traversal(root *TreeNode, res []int) []int {
	if root == nil {
		return res
	}
	res = traversal(root.Left, res)
	res = append(res, root.Val)
	return traversal(root.Right, res)
}

/*
*
中序遍历就是 左 中 右
*/
func inorderTraversalV1(root *TreeNode) []int {
	return traversal(root, nil)
}

/*
*
迭代法
原理 左边全部处理了,当前值插入,处理右边,右边处理完后 相当于 上一层的左边全部处理了 再进上一层处理
一直取左 如果没有左了 此值就是最小值,再进此值的右边取左
*/
func inorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 8)
	stack := make([]*TreeNode, 0, 8) // 存放有左子节点的节点
	curr := root
	for len(stack) > 0 || curr != nil {
		if curr != nil { // 如果当前有未扫描的节点 加入栈中,扫描其左节点
			stack = append(stack, curr)
			curr = curr.Left
		} else { // 如果左边没有元素了 取当前中节点 并扫描右节点
			item := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, item.Val)
			curr = item.Right
		}
	}
	return res
}

/*
莫里斯遍历 空间复杂度O(1) 但是会改变原节点
将当前节点移动到左节点的最右节点下
*/
func inorderTraversalV3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0, 8)
	for curr := root; curr != nil; { // 将当前节点挂载到左子节点的最右节点上
		if curr.Left == nil { // 如果没有左子 此值就是最小 插入结果 并从右边继续
			res = append(res, curr.Val)
			curr = curr.Right
			continue
		}
		// 如果有左 找到左子的最右节点
		left, rightEnd := curr.Left, curr.Left
		for left.Right != nil {
			rightEnd = left.Right
			left = rightEnd
		}
		left = curr.Left
		rightEnd.Right = curr // 将当前节点挂载的左子的最右节点上
		curr.Left = nil       // 与当前左子分离
		curr = left           // 将左子继续向下挂载
	}
	return res
}

var inorderTraversal = inorderTraversalV3

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(inorderTraversal(root))
}
