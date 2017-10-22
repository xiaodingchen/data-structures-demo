package stack

import "errors"

// Node 节点结构
type Node struct {
	data Elem
	next *Node
}

// StackLink链接栈
type StackLink struct {
	top *Node
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
}

// Pop 弹出一个元素
func (s *StackLink) Pop(e *Elem) error {
	if s.Empty() {
		return errors.New("empty stack")
	}

	*e = s.top.data
	node := s.top
	s.top = node.next
	// 释放内存空间

	return nil
}

// GetTop 获得栈顶元素
func (s *StackLink) GetTop(node *Node) {
	node = s.top
}

// Empty 是否为空栈
func (s *StackLink) Empty() bool {
	return s.top == nil
}
