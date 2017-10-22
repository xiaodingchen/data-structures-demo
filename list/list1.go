package list

type Student struct {
	Name   string
	Age    int
	Source int
}

type ListStu struct {
	stu  Student
	next *ListStu
}

// NewListStu 创建一个头结点
func NewListStu() *ListStu {
	listStu := ListStu{}
	listStu.next = nil
	return &listStu
}

func ListAdd(head *ListStu, stu Student) {
	p := head
	for p.next != nil {
		p = p.next
	}
	var node ListStu
	p.next = &node
	p.stu = stu

}

// ListInsert 插入一个元素
func ListInsert(head *ListStu, i int, stu Student) {
	p := head
	var j int
	for j = 0; j < i-1; j++ {
		p = p.next
	}
	// p 的位置为 i-1
	// 将i-1的节点的next指向i，即插入成功,新插入的节点next指向原来p的next值
	var node ListStu
	node.stu = stu
	node.next = p.next
	p.next = &node
}

// GetElem 获取指定的链表
func GetElem(head *ListStu, i int) *Student {
	p := head
	var j int

	for j = 0; j < i; j++ {
		p = p.next
	}

	if p == nil {
		return nil
	}
	return &(p.stu)

}

// ListLength 获取链表长度
func ListLength(head *ListStu) int {
	l := 0
	p := head
	for p.next != nil {
		l++
		p = p.next
	}

	return l
}

// IsEmpty 判断链表是否为空
func IsEmpty(head *ListStu) bool {
	return head.next == nil
}

// ListDelete 删除一个节点
func ListDelete(head *ListStu, i int) *Student {
	p := head
	for j := 0; j < i-1; j++ {
		p = p.next
	}

	// 当前p为i-1
	if p == nil {
		return nil
	}

	// 获取i的节点
	q := p.next
	// i-1 的next指向i的next
	p.next = q.next

	return &(q.stu)
}
