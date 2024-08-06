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
广度优先算法
进入第i层 加入结果 将第i+1层加入待遍历区 直到待遍历区为空
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 8)
	stack := []*TreeNode{root} // 一层节点
	for len(stack) > 0 {       // 待遍历区
		arr := make([]int, len(stack))
		tempStack := make([]*TreeNode, 0, len(stack)*2)
		for i, node := range stack {
			arr[i] = node.Val
			if node.Left != nil {
				tempStack = append(tempStack, node.Left)
			}
			if node.Right != nil {
				tempStack = append(tempStack, node.Right)
			}
		}
		res = append(res, arr)
		stack = tempStack
	}
	return res
}
func main() {
	root := &TreeNode{Val: 1,
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
	fmt.Println(levelOrder(root))
}
