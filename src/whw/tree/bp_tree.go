package tree

import (
	"fmt"
	"strings"
)

type btpNodeType int

const (
	Index btpNodeType = 0
	Leaf  btpNodeType = 1
)

type bptNode struct {
	parent   *bptNode
	key      []int
	val      []interface{}
	curNum   int
	nodeType btpNodeType
	right    *bptNode // 叶子节点链表
}

type BPlusTree struct {
	root *bptNode
	num  int
}

func NewBpt(n int) *BPlusTree {
	return &BPlusTree{
		root: nil,
		num:  n,
	}
}

func (n *bptNode) isFull() bool {
	return n.curNum == len(n.key)
}

func (n *bptNode) isLeaf() bool {
	return n.nodeType == Leaf
}

func (n *bptNode) isRoot() bool {
	return n.parent == nil
}

func (bpt *BPlusTree) Print() {
	leftLeaf := bpt.root
	for {
		if leftLeaf.isLeaf() {
			break
		}
		leftLeaf = leftLeaf.val[0].(*bptNode)
	}

	box := make([]string, 0)
	for leftLeaf != nil {
		s := ""
		for i := 0; i < leftLeaf.curNum; i++ {
			s += fmt.Sprint(leftLeaf.val[i], " ")
		}
		leftLeaf = leftLeaf.right
		box = append(box, strings.TrimSpace(s))
	}
	for i := range box {
		fmt.Printf("(%s)", box[i])
		if i != len(box)-1 {
			fmt.Print(" -> ")
		}
	}

	fmt.Println()
}

func (bpt *BPlusTree) newEmptyNode(nodeType btpNodeType) *bptNode {
	valLen := bpt.num
	if nodeType == Index {
		valLen += 1
	}
	return &bptNode{
		parent:   nil,
		key:      make([]int, bpt.num),
		val:      make([]interface{}, valLen),
		curNum:   0,
		nodeType: nodeType,
		right:    nil,
	}
}

func (bpt *BPlusTree) insertToIndexNode(p, n *bptNode, key int) {
	i := 0
	for ; i < p.curNum; i++ {
		if key <= p.key[i] {
			break
		}
	}

	if i < p.curNum {
		copy(p.key[i+1:], p.key[i:])
		copy(p.val[i+2:], p.val[i+1:]) // 注意：索引节点下存的数据是下一级节点的指针，数量为key的数量+1
	}

	p.key[i] = key
	p.val[i+1] = n
	p.curNum++
	n.parent = p

	if !p.isFull() {
		return
	}

	nn, newKey := bpt.divideIndexNode(p) // 索引节点分裂
	if p.isRoot() {
		bpt.newParentRootNode(p, nn, newKey)
	} else {
		bpt.insertToIndexNode(p.parent, nn, newKey)
	}
}

func (bpt *BPlusTree) insertToLeafNode(n *bptNode, key int, val interface{}) {
	i := 0
	for ; i < n.curNum; i++ {
		if key <= n.key[i] {
			break
		}
	}

	if i < n.curNum {
		copy(n.key[i+1:], n.key[i:]) // 后段，向后移动一个位置
		copy(n.val[i+1:], n.val[i:])
	}

	n.key[i] = key // 插入数据
	n.val[i] = val
	n.curNum++

	if !n.isFull() {
		return
	}

	nn, newKey := bpt.divideLeafNode(n) // 叶子节点分裂
	if n.isRoot() {
		// 如果是根节点，当前节点分裂一个新的节点出来，然后从新节点上提一个key作为父节点出来
		bpt.newParentRootNode(n, nn, newKey)
	} else {
		// 不是根节点，当前节点分裂出一个新的节点，然后新节点插入到父节点中
		bpt.insertToIndexNode(n.parent, nn, newKey)
	}
}

func (bpt *BPlusTree) newParentRootNode(ln, rn *bptNode, key int) {
	root := bpt.newEmptyNode(Index)
	bpt.root = root

	root.key[0] = key
	root.val[0] = ln
	root.val[1] = rn
	ln.parent = root
	rn.parent = root
	root.curNum = 1
}

func (bpt *BPlusTree) divideLeafNode(curNode *bptNode) (newNode *bptNode, newKey int) {
	idx := bpt.num / 2

	newNode = bpt.newEmptyNode(Leaf)

	newNode.right = curNode.right
	curNode.right = newNode

	curNode.curNum = idx
	newNode.curNum = bpt.num - idx
	newKey = curNode.key[idx]

	// 后半段迁移到新节点
	copy(newNode.key, curNode.key[idx:])
	copy(newNode.val, curNode.val[idx:])

	// 清理被拷贝走的数据
	copy(curNode.key[idx:], make([]int, len(curNode.key)-idx))
	copy(curNode.val[idx:], make([]interface{}, len(curNode.val)-idx))

	return
}

func (bpt *BPlusTree) divideIndexNode(curNode *bptNode) (newNode *bptNode, newKey int) {
	idx := bpt.num / 2

	newKey = curNode.key[idx]
	curNode.key[idx] = 0

	newNode = bpt.newEmptyNode(Index)
	curNode.curNum = idx
	newNode.curNum = bpt.num - idx - 1

	copy(newNode.key, curNode.key[idx+1:])
	copy(newNode.val, curNode.val[idx+1:])

	copy(curNode.key[idx+1:], make([]int, len(curNode.key)-idx-1))
	copy(curNode.val[idx+1:], make([]interface{}, len(curNode.val)-idx-1))

	return
}

func (bpt *BPlusTree) findLeafNodeToInsert(n *bptNode, key int) *bptNode {
	if n.isLeaf() {
		return n
	}

	i := 0
	for ; i < n.curNum; i++ {
		if key <= n.key[i] {
			break
		}
	}
	n = n.val[i].(*bptNode)
	return bpt.findLeafNodeToInsert(n, key)
}

func (bpt *BPlusTree) findInLeafNode(node *bptNode, key int) (val interface{}, ok bool) {
	for i := 0; i < node.curNum; i++ {
		if node.key[i] == key {
			return node.val[i], true
		}
	}
	return nil, false
}

func (bpt *BPlusTree) minNodeNum() int {
	return (bpt.num+1)/2 - 1
}

func (bpt *BPlusTree) updateIndex(node *bptNode) {
	if node.isRoot() {
		return
	}

	parent := node.parent
	i := 0
	for ; i < parent.curNum; i++ {
		if parent.val[i] == node {
			break
		}
	}

	if i == 0 {
		bpt.updateIndex(parent)
	} else {
		parent.key[i-1] = node.key[0]
	}
}

// 向前找兄弟节点，优先找可以借用元素的节点
// 为什么向前找，而不是向后找
// 向前：如果借用元素，借用末尾位置，最后一个位置清空即可
// 向后：如果借用元素，借用开始位置，剩余位置全部前移一个空位
func (bpt *BPlusTree) getSiblingNode(node *bptNode) (n *bptNode, isBefore bool) {
	p := node.parent
	i := 0
	for ; i <= p.curNum; i++ {
		if p.val[i] == node {
			break
		}
	}

	if i == 0 {
		return p.val[i+1].(*bptNode), false
	} else {
		n = p.val[i-1].(*bptNode)
		if i == p.curNum-1 || n.curNum > bpt.minNodeNum() {
			return n, true
		}
		next := p.val[i].(*bptNode)
		if next.curNum > bpt.minNodeNum() {
			return next, false
		} else {
			return n, true
		}
	}
}

func (bpt *BPlusTree) borrowLeafElement(src, dst *bptNode, last bool) {
	idx := 0 // 借第一个
	if last {
		idx = src.curNum - 1 // 借最后一个
	}

	key := src.key[idx]
	val := src.val[idx]
	src.key[idx] = 0
	src.val[idx] = nil
	src.curNum -= 1

	if !last { // 借第一个，需要所有元素前移，更新索引
		copy(src.key[0:], src.key[1:src.curNum+1])
		copy(src.val[0:], src.val[1:src.curNum+1])
	}

	bpt.insertToLeafNode(dst, key, val)

	if last {
		bpt.updateIndex(dst)
	} else {
		bpt.updateIndex(src)
	}
}

func (bpt *BPlusTree) borrowIndexElement(src, dst *bptNode, last bool) {
	var key int
	var val interface{}
	if last { // 借最后一个
		key = src.key[src.curNum-1]
		val = src.val[src.curNum]

		src.key[src.curNum-1] = 0 //
		src.val[src.curNum] = nil
	} else { // 借第一个
		key = src.key[0]
		val = src.val[0]

		copy(src.key[0:], src.key[1:src.curNum+1]) // 前移
		copy(src.val[0:], src.val[1:src.curNum+1])
	}

	bpt.insertToIndexNode(dst, val.(*bptNode), key)

	if last {
		bpt.updateIndex(dst)
	} else {
		bpt.updateIndex(src)
	}
}

func (bpt *BPlusTree) combineIndexNode(ln, rn *bptNode) {
	// 从父节点中删除
	key := bpt.deleteIndexNode(rn)
	ln.key[ln.curNum] = key
	ln.curNum += rn.curNum + 1

	// 从父节点获取索引key，添加到中间
	copy(ln.key[ln.curNum:], rn.key[:rn.curNum])
	copy(ln.val[ln.curNum:], rn.val[:rn.curNum+1])
}

func (bpt *BPlusTree) deleteIndexNode(child *bptNode) (key int) {
	parent := child.parent

	i := 0
	for ; i <= parent.curNum; i++ {
		if parent.val[i] == child {
			break
		}
	}

	key = parent.key[i-1]

	parent.key[i-1] = 0
	parent.val[i] = nil
	if i != parent.curNum {
		copy(parent.key[i-1:], parent.key[i:])
		copy(parent.val[i:], parent.val[i+1:])
	}

	parent.curNum -= 1

	if parent.isRoot() || parent.curNum >= bpt.minNodeNum() {
		return
	}

	sibling, isBefore := bpt.getSiblingNode(parent)
	if sibling.curNum > bpt.minNodeNum() {
		bpt.borrowIndexElement(sibling, parent, isBefore)
		return
	}

	if isBefore {
		bpt.combineIndexNode(sibling, parent)
	} else {
		bpt.combineIndexNode(parent, sibling)
	}

	return
}

func (bpt *BPlusTree) combineLeafNode(ln, rn *bptNode) {
	// 从父节点中删除
	bpt.deleteIndexNode(rn)

	copy(ln.key[ln.curNum:], rn.key[:rn.curNum])
	copy(ln.val[ln.curNum:], rn.val[:rn.curNum])
	ln.curNum += rn.curNum
	ln.right = rn.right

	// 如果父节点为root，并且没有key，那么当前节点替代父节点
	if ln.parent.isRoot() && ln.parent.curNum == 0 {
		bpt.root = ln
		ln.parent = nil
	}
}

func (bpt *BPlusTree) deleteInLeaf(node *bptNode, key int) (val interface{}, ok bool) {
	i := 0
	for ; i < node.curNum; i++ {
		if node.key[i] == key {
			break
		}
	}

	if i >= node.curNum {
		return nil, false
	}

	val = node.val[i]
	ok = true

	copy(node.key[i:], node.key[i+1:])
	copy(node.val[i:], node.val[i+1:])
	node.curNum--

	if node.isRoot() {
		return
	}

	// 检查节点数量是否足够
	if node.curNum >= bpt.minNodeNum() {
		if i == 0 { // 更新索引
			bpt.updateIndex(node)
		}
		return
	}

	// 不够
	// 1. 检查相邻的兄弟节点，如果兄弟节点数量大于 min，则向其借一个节点
	// 2. 如果兄弟节点数量小于等于 min，那么可以合兄弟节点合并

	// isBefore==true: 找到的兄弟节点在当前节点左边，就应该 [借最后一个元素 / 合并时sibling在前]
	sibling, isBefore := bpt.getSiblingNode(node)
	if sibling.curNum > bpt.minNodeNum() {
		bpt.borrowLeafElement(sibling, node, isBefore)
		return
	}

	// 合并
	if isBefore {
		bpt.combineLeafNode(sibling, node)
	} else {
		bpt.combineLeafNode(node, sibling)
	}

	return
}

/// --------------------------------------------------------------------------------------------------------------------
func (bpt *BPlusTree) Insert(key int, val interface{}) {
	if bpt.root == nil {
		bpt.root = bpt.newEmptyNode(Leaf)
		bpt.insertToLeafNode(bpt.root, key, val)
		return
	}

	n := bpt.findLeafNodeToInsert(bpt.root, key)
	bpt.insertToLeafNode(n, key, val)
}

func (bpt *BPlusTree) Find(key int) (val interface{}, ok bool) {
	node := bpt.root
	for node != nil {
		if node.isLeaf() {
			return bpt.findInLeafNode(node, key)
		}

		i := 0
		for ; i < node.curNum; i++ {
			if key <= node.key[i] {
				break
			}
		}
		node = node.val[i].(*bptNode)
	}
	return nil, false
}

func (bpt *BPlusTree) Delete(key int) (val interface{}, ok bool) {
	node := bpt.root
	for node != nil {
		if node.isLeaf() {
			return bpt.deleteInLeaf(node, key)
		}

		i := 0
		for ; i < node.curNum; i++ {
			if key < node.key[i] {
				break
			}
		}
		node, _ = node.val[i].(*bptNode)
	}
	return nil, false
}
