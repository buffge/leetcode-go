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
先序遍历 将右节点放左节点前面  依次取 当前 左 右
*/
func flattenV1(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode // 上一个节点 用于连接
	for len(stack) > 0 {
		node := stack[len(stack)-1] // 移出当前节点
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, node // 连接
		}
		if node.Right != nil { // 先放入右
			stack = append(stack, node.Right)
		}
		if node.Left != nil { // 再放左
			stack = append(stack, node.Left)
		}
		prev = node // 更新链表的上一个
	}
}

/*
寻找前驱节点 空间复杂度O(1)
先序遍历 定理 中 左 右
当前节点遍历的下一个节点一定是左子节点
当前右子节点的前一个遍历节点移动是左子树中最右节点
所以要将左子插入到右子节点下 再将原右子树挂载到左子树最右节点下

遍历根节点
1.当根节点左边不为空时 将左边插入到右边
2.当左边为空 根节点设置为当前节点的右子节点
*/
func flattenV2(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil { // 将左边依次挂到右边上 挂完后当前节点的左就是空
			left := curr.Left
			endNode := left // 找到当前节点左树中最右节点endNode  此endNode的下一个节点就是当前节点的右子节点
			for endNode.Right != nil {
				endNode = endNode.Right
			}
			endNode.Right = curr.Right        // 将右子树挂载左子最后一个节点后 即 先序遍历顺序
			curr.Left, curr.Right = nil, left // 当前节点计算完毕 计算当前节点的左子树
		}
		curr = curr.Right // 遍历右节点
	}
}

var flatten = flattenV2

func main() {
	root := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 5,
			Right: &TreeNode{Val: 6,
				Right: nil},
		},
	}
	flatten(root)
	fmt.Println(root)

}
