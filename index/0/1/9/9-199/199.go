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
层序遍历 取每一层的最后一个
*/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var res []int
	for len(stack) > 0 {
		endNode := stack[len(stack)-1]
		res = append(res, endNode.Val)
		tempStack := make([]*TreeNode, 0, len(stack))
		for _, node := range stack {
			if node.Left != nil {
				tempStack = append(tempStack, node.Left)
			}
			if node.Right != nil {
				tempStack = append(tempStack, node.Right)
			}
		}
		stack = tempStack
	}
	return res
}
func main() {
	root := &TreeNode{
		Val: 32,
		Left: &TreeNode{
			Val: 26,
			Left: &TreeNode{
				Val:   19,
				Right: &TreeNode{Val: 27},
			},
		},
		Right: &TreeNode{
			Val: 47,
			//Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 56},
		},
	}
	fmt.Println(rightSideView(root))
}
