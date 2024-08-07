package utils

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	res := &TreeNode{Val: *vals[0]}
	stack := make([]*TreeNode, 0, len(vals))
	stack = append(stack, res)

	for len(stack) > 0 && len(vals) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(vals) > 0 {
			val := vals[0]
			vals = vals[1:]
			if val != nil {
				curr.Left = &TreeNode{Val: *val}
				stack = append(stack, curr.Left)
			}
		}
		if len(vals) > 0 {
			val := vals[0]
			vals = vals[1:]
			if val != nil {
				curr.Right = &TreeNode{Val: *val}
				stack = append(stack, curr.Right)
			}
		}
	}
	return res
}

/*
*
前序遍历 中 左 右
Preorder Traversal
*/
func preOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var res []int
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}

/*
*
中序遍历 左 中 右
*/
func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var stack []*TreeNode
	var prev *TreeNode
	curr := root
	var res []int
	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil && node.Right != prev {
			stack = append(stack, node.Right)
		}
		prev = node
	}
	return res
}

/*
*
层序遍历
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	var res [][]int
	for len(stack) > 0 {
		tempStack := make([]*TreeNode, 0, len(stack)>>1)
		arr := make([]int, len(stack))
		for i, v := range stack {
			if v.Left != nil {
				tempStack = append(tempStack, v.Left)
			}
			if v.Right != nil {
				tempStack = append(tempStack, v.Right)
			}
			arr[i] = v.Val
		}
		res = append(res, arr)
		stack = tempStack
	}
	return res
}
