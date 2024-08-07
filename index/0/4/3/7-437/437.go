package main

import (
	"fmt"
	"github.com/buffge/leetcode-go/utils"
)

type (
	TreeNode = utils.TreeNode
)

func dfs(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Val == targetSum {
		res++
	}
	res += dfs(root.Left, targetSum-root.Val)
	res += dfs(root.Right, targetSum-root.Val)
	return res
}

/*
todo 迭代版 前缀版
dfs
计算路径 如果root==sum, +1,如果 root+左子==sum+1 如果 root+右子==sum+1
*/
func pathSumV1(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := dfs(root, targetSum)
	res += pathSum(root.Left, targetSum) // 递归 允许下级任意一个节点作为起始节点
	res += pathSum(root.Right, targetSum)
	return res
}

var pathSum = pathSumV1

func main() {
	root := &TreeNode{Val: 10,
		Left: &TreeNode{Val: 5,
			Left: &TreeNode{Val: 3,
				Left:  &TreeNode{Val: 3},
				Right: &TreeNode{Val: -2},
			},
			Right: &TreeNode{Val: 2,
				Right: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{Val: -3,
			Right: &TreeNode{Val: 11,
				Right: nil},
		},
	}
	targetSum := 8
	fmt.Println(pathSum(root, targetSum))
}
