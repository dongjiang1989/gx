///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Nov 27 2016 10:53:19]
// Generate from:
//   [github.com/vipally/gx/stl/gp/tree.gp]
//   [github.com/vipally/gx/fs/tree.gpg] [tree_filehash]
//
// Tool [github.com/vipally/gogp] info:
// CopyRight 2016 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Blog    : http://blog.csdn.net/vipally
// Site    : https://github.com/vipally
// BuildAt : [Oct 24 2016 20:25:45]
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

//this file defines a template tree structure just like file system

package fs

////////////////////////////////////////////////////////////////////////////////

//tree strture
type FileHashTree struct {
	root *FileHashTreeNode
}

//new container
func NewFileHashTree() *FileHashTree {
	p := &FileHashTree{}
	return p
}

//tree node
type FileHashTreeNode struct {
	val      FileHash
	children FileHashTreeNodeSortSlice
}

func (this *FileHashTreeNode) Less(right *FileHashTreeNode) (ok bool) {

	ok = this.val.Less(right.val)

	return
}

func (this *FileHashTreeNode) SortChildren() {
	this.children.Sort()
}

func (this *FileHashTreeNode) Children() []*FileHashTreeNode {
	return this.children.Buffer()
}

//add a child
func (this *FileHashTreeNode) AddChild(v FileHash, idx int) (child *FileHashTreeNode) {
	n := &FileHashTreeNode{val: v}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *FileHashTreeNode) AddChildNode(node *FileHashTreeNode, idx int) (child *FileHashTreeNode) {
	this.children.Insert(node, idx)
	return node
}

//cound of children
func (this *FileHashTreeNode) NumChildren() int {
	return this.children.Len()
}

//get child
func (this *FileHashTreeNode) GetChild(idx int) (child *FileHashTreeNode, ok bool) {
	child, ok = this.children.Get(idx)
	return
}

//remove child
func (this *FileHashTreeNode) RemoveChild(idx int) (child *FileHashTreeNode, ok bool) {
	child, ok = this.children.Remove(idx)
	return
}

//create a visitor
func (this *FileHashTreeNode) Visitor() (v *FileHashTreeNodeVisitor) {
	v = &FileHashTreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *FileHashTreeNode) All() (list []FileHash) {
	list = append(list, this.val)
	for _, v := range this.children.Buffer() {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type FileHashTreeNodeVisitor struct {
	node         *FileHashTreeNode
	parents      []*FileHashTreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *FileHashTreeNodeVisitor) push(n *FileHashTreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *FileHashTreeNodeVisitor) pop() (n *FileHashTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *FileHashTreeNodeVisitor) tail() (n *FileHashTreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *FileHashTreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *FileHashTreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *FileHashTreeNodeVisitor) top_right(n *FileHashTreeNode) (p *FileHashTreeNode) {
	if n != nil {
		l := n.children.Len()
		for l > 0 {
			this.push(n, l-1)
			n = n.children.MustGet(l - 1)
			l = n.children.Len()
		}
		p = n
	}
	return
}

//visit next node
func (this *FileHashTreeNodeVisitor) Next() (ok bool) {
	if this.node != nil { //check if has any children
		if this.node.children.Len() > 0 {
			this.push(this.node, 0)
			this.node = this.node.children.MustGet(0)
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < p.children.Len()-1 { //next brother
			bIdx++
			this.node = p.children.MustGet(bIdx)
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//visit previous node
func (this *FileHashTreeNodeVisitor) Prev() (ok bool) {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		if ok = this.node != nil; ok {
			//do nothing
		}
		return
	}

	if this.node != nil { //check if has any children
		p, bIdx := this.tail()
		if bIdx > 0 {
			bIdx--
			this.update_tail(bIdx)
			this.node = this.top_right(p.children.MustGet(bIdx))
		} else {
			this.node = p
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		//do nothing
	}
	return
}

//get node data
func (this *FileHashTreeNodeVisitor) Get() (data FileHash) {
	if nil != this.node {
		data = this.node.val
	}
	return
}
