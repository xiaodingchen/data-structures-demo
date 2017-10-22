package list

import (
	"testing"
)

func TestLinkListStart(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		headNode := NewListStu()
		stu1 := Student{Name: "jack", Age: 18, Source: 98}
		ListAdd(headNode, stu1)
		t.Logf("linklist is empty: %v", IsEmpty(headNode))
		stu2 := Student{Name: "marry", Age: 18, Source: 88}
		ListAdd(headNode, stu2)
		t.Logf("elem: %v", GetElem(headNode, 1))
		t.Logf("listling len: %d", ListLength(headNode))
		t.Logf("node: %v", headNode.stu)
	})
}
