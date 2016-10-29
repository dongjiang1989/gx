//#GOGP_IGNORE_BEGIN
///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/vipally/gogp]
// Last update at: [Sun Oct 30 2016 04:09:50]
// Generate from:
//   [github.com/vipally/gx/stl/gp/tree.go]
//   [github.com/vipally/gx/stl/gp/gp.gpg] [GOGP_REVERSE_tree]
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
//#GOGP_IGNORE_END

//this file defines a template tree structure just like file system

<PACKAGE>

import (
	"sort"
)

//tree strture
type <CONTAINER_NAME_PREFIX>Tree struct {
	root *<CONTAINER_NAME_PREFIX>TreeNode
}

//new container
func New<CONTAINER_NAME_PREFIX>Tree() *<CONTAINER_NAME_PREFIX>Tree {
	return &<CONTAINER_NAME_PREFIX>Tree{}
}

//tree node
type <CONTAINER_NAME_PREFIX>TreeNode struct {
	<VALUE_TYPE>
	children __<CONTAINER_NAME_PREFIX>TreeNodeSortSlice
}

func (this *<CONTAINER_NAME_PREFIX>TreeNode) SortChildren() {
	this.children.Sort()
}

func (this *<CONTAINER_NAME_PREFIX>TreeNode) Children() []*<CONTAINER_NAME_PREFIX>TreeNode {
	return this.children.Buffer()
}

//add a child
func (this *<CONTAINER_NAME_PREFIX>TreeNode) AddChild(v <VALUE_TYPE>, idx int) (child *<CONTAINER_NAME_PREFIX>TreeNode) {
	n := &<CONTAINER_NAME_PREFIX>TreeNode{<VALUE_TYPE>: v, children: nil}
	return this.AddChildNode(n, idx)
}

//add a child node
func (this *<CONTAINER_NAME_PREFIX>TreeNode) AddChildNode(node *<CONTAINER_NAME_PREFIX>TreeNode, idx int) (child *<CONTAINER_NAME_PREFIX>TreeNode) {
	if idx >= 0 && idx < len(this.children) {
		right := this.children[idx+1:]
		this.children = append(this.children[:idx], node)
		this.children = append(this.children, right...)
	} else {
		this.children = append(this.children, node)
	}
	return node
}

//cound of children
func (this *<CONTAINER_NAME_PREFIX>TreeNode) NumChildren() int {
	return len(this.children)
}

//get child
func (this *<CONTAINER_NAME_PREFIX>TreeNode) GetChild(idx int) (child *<CONTAINER_NAME_PREFIX>TreeNode, ok bool) {
	if ok = idx >= 0 && idx < len(this.children); ok {
		child = this.children[idx]
	}
	return
}

//remove child
func (this *<CONTAINER_NAME_PREFIX>TreeNode) RemoveChild(idx int) (child *<CONTAINER_NAME_PREFIX>TreeNode, ok bool) {
	if child, ok = this.GetChild(idx); ok {
		this.children = append(this.children[:idx], this.children[idx+1:]...)
	}
	return
}

//create a visitor
func (this *<CONTAINER_NAME_PREFIX>TreeNode) Visitor() (v *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) {
	v = &<CONTAINER_NAME_PREFIX>TreeNodeVisitor{}
	v.push(this, -1)
	return
}

//get all node data
func (this *<CONTAINER_NAME_PREFIX>TreeNode) All() (list []<VALUE_TYPE>) {
	list = append(list, this.<VALUE_TYPE>)
	for _, v := range this.children {
		list = append(list, v.All()...)
	}
	return
}

//tree node visitor
type <CONTAINER_NAME_PREFIX>TreeNodeVisitor struct {
	node         *<CONTAINER_NAME_PREFIX>TreeNode
	parents      []*<CONTAINER_NAME_PREFIX>TreeNode
	brotherIdxes []int
	//visit order: this->child->brother
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) push(n *<CONTAINER_NAME_PREFIX>TreeNode, bIdx int) {
	this.parents = append(this.parents, n)
	this.brotherIdxes = append(this.brotherIdxes, bIdx)
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) pop() (n *<CONTAINER_NAME_PREFIX>TreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n, bIdx = this.tail()
		this.parents = this.parents[:l-1]
		this.brotherIdxes = this.brotherIdxes[:l-1]
	}
	return
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) tail() (n *<CONTAINER_NAME_PREFIX>TreeNode, bIdx int) {
	l := len(this.parents)
	if l > 0 {
		n = this.parents[l-1]
		bIdx = this.brotherIdxes[l-1]
	}
	return
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) depth() int {
	return len(this.parents)
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) update_tail(bIdx int) bool {
	l := len(this.parents)
	if l > 0 {
		this.brotherIdxes[l-1] = bIdx
		return true
	}
	return false
}

func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) top_right(n *<CONTAINER_NAME_PREFIX>TreeNode) (p *<CONTAINER_NAME_PREFIX>TreeNode) {
	if n != nil {
		l := len(n.children)
		for l > 0 {
			this.push(n, l-1)
			n = n.children[l-1]
			l = len(n.children)
		}
		p = n
	}
	return
}

//visit next node
func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) Next() (data *<VALUE_TYPE>, ok bool) {
	if this.node != nil { //check if has any children
		if len(this.node.children) > 0 {
			this.push(this.node, 0)
			this.node = this.node.children[0]
		} else {
			this.node = nil
		}
	}
	for this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, bIdx := this.tail()
		if bIdx < 0 { //ref parent
			this.node = p
			this.pop()
		} else if bIdx < len(p.children)-1 { //next brother
			bIdx++
			this.node = p.children[bIdx]
			this.update_tail(bIdx)
		} else { //no more brothers
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		data = this.Get()
	}
	return
}

//visit previous node
func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) Prev() (data *<VALUE_TYPE>, ok bool) {
	if this.node == nil && this.depth() > 0 { //check if has any brothers or uncles
		p, _ := this.pop()
		this.node = this.top_right(p)
		if ok = this.node != nil; ok {
			data = this.Get()
		}
		return
	}

	if this.node != nil { //check if has any children
		p, bIdx := this.tail()
		if bIdx > 0 {
			bIdx--
			this.update_tail(bIdx)
			this.node = this.top_right(p.children[bIdx])
		} else {
			this.node = p
			this.pop()
		}
	}
	if ok = this.node != nil; ok {
		data = this.Get()
	}
	return
}

//get node data
func (this *<CONTAINER_NAME_PREFIX>TreeNodeVisitor) Get() *<VALUE_TYPE> {
	return &this.node.<VALUE_TYPE>
}

//for sort
type __<CONTAINER_NAME_PREFIX>TreeNodeSortSlice []*<CONTAINER_NAME_PREFIX>TreeNode

func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Sort() {
	sort.Sort(this)
}

//data
func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Buffer() []*<CONTAINER_NAME_PREFIX>TreeNode {
	return *this
}

//push
func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Push(v *<CONTAINER_NAME_PREFIX>TreeNode) int {
	*this = append(*this, v)
	return this.Len()
}

func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Pop() (r *<CONTAINER_NAME_PREFIX>TreeNode) {
	if len(*this) > 0 {
		r = (*this)[len(*this)-1]
	}
	*this = (*this)[:len(*this)-1]
	return
}

//len
func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Len() int {
	return len(this.Buffer())
}

//sort by Hash decend,the larger one first
func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Less(i, j int) (ok bool) {
	l, r := (*this)[i], (*this)[j]

	//#GOGP_IFDEF GOGP_HasLess
	ok = l.Less(r.<VALUE_TYPE>)
	//#GOGP_ELSE
	ok = l.<VALUE_TYPE> < r.<VALUE_TYPE>
	//#GOGP_ENDIF

	return
}

//swap
func (this *__<CONTAINER_NAME_PREFIX>TreeNodeSortSlice) Swap(i, j int) {
	(*this)[i], (*this)[j] = (*this)[j], (*this)[i]
}

