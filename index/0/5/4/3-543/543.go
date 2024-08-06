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

深度优先算法 迭代版
非原地  偷懒将val存深度了 可以用hash 也可以将节点包一层
计算"经过"某个节点的最大路径就是 左边深度+右边深度 + 1或0,  自身有子节点时为1
*/
func diameterOfBinaryTreeV1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxVal := 0
	var stack []*TreeNode // 存放待计算的节点
	calcRoot := root      // 待计算树的根节点 需将此节点和此节点的所有左子节点加入待计算中
	var prev *TreeNode    // 上一轮计算的节点
	for len(stack) > 0 || calcRoot != nil {
		for calcRoot != nil { // 当搜索节点的左子节点全部加入栈中
			stack = append(stack, calcRoot)
			calcRoot = calcRoot.Left
		}
		node := stack[len(stack)-1]                  // 节点的左子节点已计算完毕
		if node.Right == nil || node.Right == prev { // 如果没有右节点或者右节点是上一轮计算过的 更新当前节点值并移除
			// 此时 node 的左右均已计算完毕
			if node.Left == nil && node.Right == nil {
				node.Val = 1 // 给父节点深度+1
			} else { // 如果当前计算节点有子级 更新最大直径还当前节点的深度
				leftDepth, rightDepth := 0, 0
				if node.Left != nil {
					leftDepth = node.Left.Val
				}
				if node.Right != nil {
					rightDepth = node.Right.Val
				}
				maxVal = max(maxVal, leftDepth+rightDepth)
				node.Val = max(leftDepth, rightDepth) + 1
			}
			stack = stack[:len(stack)-1] // 此节点计算完毕 移除
			prev = node                  // 更新上一轮计算的节点
		} else { // 如果有未计算的右子节点 需将所有的左子节点加入待计算中
			calcRoot = node.Right
		}
	}
	return maxVal
}

/*
*
todo 默写一遍
*/
func diameterOfBinaryTreeV2(root *TreeNode) int {
	return 0
}

var diameterOfBinaryTree = diameterOfBinaryTreeV1

func newInt(val int) *int {
	return &val
}
func main() {
	nodes := []*int{
		newInt(1), newInt(1), newInt(1),
		nil, nil,
		newInt(1), newInt(1), newInt(1), newInt(1), newInt(1),
		nil,
		newInt(1),
		nil,
		newInt(1), newInt(1),
		nil, nil,
		newInt(1), newInt(1), newInt(1),
		nil,
		newInt(1),
		nil, nil,
		newInt(1), newInt(1),
		nil, nil, nil,
		newInt(1),
	}
	root := utils.GenerateTree(nodes)
	root = &TreeNode{Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	root = &TreeNode{Val: 1}
	fmt.Println(diameterOfBinaryTree(root))
}
