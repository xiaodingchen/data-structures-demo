package stack

const INIT_CAP = 10

// const INCRE_MENT = 5

type Elem int

type Stack struct {
	top  int
	data [INIT_CAP]Elem
}

// InitStack 创建一个空栈
func InitStack(s *Stack) {
	s.top = 0
}

// Push 增加一个栈元素
func Push(s *Stack, e Elem) bool {
	if s.top > INIT_CAP {
		return false
	}
	s.data[s.top] = e
	s.top++
	return true
}

// Pop 弹出一个元素
func Pop(s *Stack, e *Elem) bool {
	if s.top == 0 {
		return false
	}

	*e = s.data[s.top-1]
	s.top--
	return true
}

// Empty 是否是空栈
func Empty(s Stack) bool {
	return s.top == 0
}
