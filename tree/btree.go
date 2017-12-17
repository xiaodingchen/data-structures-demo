package tree

import (
	"math"
)

// TreeNode 定义树节点
type TreeNode struct {
	data   interface{}
	lchild *TreeNode
	rchild *TreeNode
	parent *TreeNode
	height int // 以该节点为根的子树高度
	size   int // 该节点子孙数，包括自身
}

// binaryTree 定义二叉树
type binaryTree struct {
	root   *TreeNode // 根节点
	height int       // 高度
	size   int       // 节点数
}

// NewBiTree 新建一个二叉树
func NewBiTree(root *TreeNode) *binaryTree {
	return &binaryTree{root: root}
}

// GetSize 获取树的节点总数
func (b *binaryTree) GetSize() int {
	return b.root.size
}

// GetHeight 获取树的高度
func (b *binaryTree) GetHeight() int {
	return b.root.height
}

// Empty 判断是否为空
func (b *binaryTree) Empty() bool {
	return b.root == nil
}

// Root 获取树的根节点
func (b *binaryTree) Root() *TreeNode {
	return b.root
}

// 获取第一个与数据e相等的节点
func (b *binaryTree) Find(e interface{}) *TreeNode {
	if b.Empty() {
		return nil
	}
	// 从根节点开始查找
	p := b.root
	// 递归查找
	return isEqual(e, p)
}

func isEqual(e interface{}, node *TreeNode) *TreeNode {
	if e == node.GetData() {
		return node
	}
	if node.HasLChild() {
		lp := isEqual(e, node.GetLChild())
		if lp != nil {
			return lp
		}
	}

	if node.HasRChild() {
		rp := isEqual(e, node.GetRChild())
		if rp != nil {
			return rp
		}
	}

	return nil

}

// NewBiTreeNode 新建一个树节点
func NewBiTreeNode(e interface{}) *TreeNode {
	return &TreeNode{data: e}
}

// GetData 获取节点数据
func (t *TreeNode) GetData() interface{} {
	return t.data
}

// SetData 设置节点数据
func (t *TreeNode) SetData(e interface{}) {
	t.data = e
}

// HasParent 判断是否有父节点
func (t *TreeNode) HasParent() bool {
	return t.parent != nil
}

// GetParent 获取父节点
func (t *TreeNode) GetParent() *TreeNode {
	return t.parent
}

// setParent 设置父节点
func (t *TreeNode) setParent(p *TreeNode) {
	t.parent = p
}

// IsLChild 判断自身是否是左孩子节点
func (t *TreeNode) IsLChild() bool {
	return t.HasParent() && t == t.parent.lchild
}

// IsRChild 判断自身是否是右孩子节点
func (t *TreeNode) IsRChild() bool {
	return t.HasParent() && t == t.parent.rchild
}

// HasLChild 是否有左孩子
func (t *TreeNode) HasLChild() bool {
	return t.lchild != nil
}

// HasRChild 是否有右孩子
func (t *TreeNode) HasRChild() bool {
	return t.rchild != nil
}

// GetLChild 获取左孩子
func (t *TreeNode) GetLChild() *TreeNode {
	return t.lchild
}

// GetRChild 获取右孩子
func (t *TreeNode) GetRChild() *TreeNode {
	return t.rchild
}

// SetLChild 设置左孩子
func (t *TreeNode) SetLChild(l *TreeNode) *TreeNode {
	oldL := t.lchild
	// 让原先的左孩子切断与父节点的联系
	if t.HasLChild() {
		t.lchild.CutParent()
	}
	// 切断新左孩子与其原父节点的联系
	if l != nil {
		l.CutParent()
		// 关联父子节点
		t.lchild = l
		l.setParent(t)
		// 重新计算当前节点的高度和容量
		t.SetHeight()
		t.SetSize()
	}

	return oldL
}

// SetRChild 设置右孩子
func (t *TreeNode) SetRChild(r *TreeNode) *TreeNode {
	oldR := t.rchild
	if t.HasRChild() {
		t.rchild.CutParent()
	}

	if r != nil {
		r.CutParent()
		t.rchild = r
		r.setParent(t)
		t.SetHeight()
		t.SetSize()
	}

	return oldR
}

// IsLeaf 判断是否为叶子节点
func (t *TreeNode) IsLeaf() bool {
	return !t.IsLChild() && !t.IsRChild()
}

// CutParent 断开与父节点的关系
func (t *TreeNode) CutParent() {
	if !t.HasParent() {
		return
	}

	if t.IsLChild() {
		// 断开该节点与父节点的联系
		t.parent.lchild = nil
	}

	if t.IsRChild() {
		// 断开该节点与父节点的联系
		t.parent.rchild = nil
	}

	// 重新计算当前节点的父节点的高度和长度
	t.parent.SetSize()
	t.parent.SetHeight()
	// 断开该节点与父节点的联系
	t.parent = nil

}

// SetSize 设置节点的长度
func (t *TreeNode) SetSize() {
	t.size = 1
	if t.HasLChild() {
		t.size += t.GetLChild().GetSize()
	}

	if t.HasRChild() {
		t.size += t.GetRChild().GetSize()
	}

	// 递归更新祖先的规模
	if t.HasParent() {
		t.parent.SetSize()
	}
}

// SetHeight 设置节点的高度
func (t *TreeNode) SetHeight() {
	newH := 0 // 新高度初始化为0，高度等于左右子树高度加1中的最大者
	if t.HasLChild() {
		newH = int(math.Max(float64(newH), float64(t.GetLChild().GetHeight())))
	}

	if t.HasRChild() {
		newH = int(math.Max(float64(newH), float64(t.GetRChild().GetHeight())))
	}

	if newH == t.height {
		return
	}

	t.height = newH
	if t.HasParent() {
		// 递归设置祖先的高度
		t.GetParent().SetHeight()
	}

}

// GetSize 获取节点的长度
func (t *TreeNode) GetSize() int {
	return t.size
}

// GetHeight 获取节点的高度
func (t *TreeNode) GetHeight() int {
	return t.height
}
