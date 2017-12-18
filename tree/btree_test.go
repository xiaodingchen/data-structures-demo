package tree

import (
	"testing"
)

func TestStart(t *testing.T) {
	root1 := NewBiTreeNode(1)
	root1.SetLChild(NewBiTreeNode(2))
	root1.SetRChild(NewBiTreeNode(3))
	root1.GetLChild().SetLChild(NewBiTreeNode(4))
	root1.GetLChild().SetRChild(NewBiTreeNode(5))
	root1.GetRChild().SetLChild(NewBiTreeNode(6))
	node2 := root1.GetLChild()
	node5 := root1.GetLChild().GetRChild()
	t.Logf("节点2是叶子节点吗？ %v", node2.IsLeaf())
	t.Logf("节点5是叶子节点吗？ %v", node5.IsLeaf())
	tree1 := NewBiTree(root1)
	t.Logf("这棵树的节点总数：%v", tree1.GetSize())
	t.Logf("这棵树的高度：%v", tree1.GetHeight())

	// 遍历
	// 中序遍历
	l := tree1.InOrder()
	for e := l.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*TreeNode)
		t.Logf("中序遍历data: %v", *obj)
	}

	// 先序遍历
	pre := tree1.PreOrder()
	for e := pre.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*TreeNode)
		t.Logf("先序遍历data: %v", *obj)
	}

	// 后序遍历
	post := tree1.PostOrder()
	for e := post.Front(); e != nil; e = e.Next() {
		obj, _ := e.Value.(*TreeNode)
		t.Logf("后序遍历data: %v", *obj)
	}
}
