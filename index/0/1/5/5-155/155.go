package main

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	if len(s.minStack) == 0 || val < s.stack[s.minStack[len(s.minStack)-1]] {
		s.minStack = append(s.minStack, len(s.stack)-1)
	} else {
		s.minStack = append(s.minStack, s.minStack[len(s.minStack)-1])
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.stack[s.minStack[len(s.minStack)-1]]
}

/**
2个栈 一个栈 存值 一个栈存插入当前值时最小的值索引
*/

func main() {
	obj := Constructor()
	obj.Push(1)
	obj.Pop()

}
