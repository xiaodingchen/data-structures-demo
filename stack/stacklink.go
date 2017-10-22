package stack

import "errors"

// Node 节点结构
type Node struct {
	data Elem
	next *Node
}

// StackLink链接栈
type StackLink struct {
	top    *Node
	length int
}

// InitStack 初始化一个栈，栈先进后出
func (s *StackLink) InitStack() {
	s.top = nil
}

// Push 添加一个栈元素
func (s *StackLink) Push(data Elem) {
	// 创造一个节点
	node := new(Node)
	node.data = data
	node.next = s.top
	s.top = node
	s.length++
}

// Pop 弹出一个元素
func (s *StackLink) Pop(e *Elem) error {
	if s.Empty() {
		return errors.New("empty stack")
	}

	*e = s.top.data
	node := s.top
	s.top = node.next
	s.length--

	return nil
}

// Empty 是否为空栈
func (s *StackLink) Empty() bool {
	return s.top == nil
}

// Length 栈的元素个数
func (s *StackLink) Length() int {
	return s.length
}
