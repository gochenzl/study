package bplustree

// B+tree

type Comparer func(interface{}, interface{}) int

type BplusTree struct {
	cmp  Comparer
	m    int
	l    int
	size int
	root commonNode
}

// m is order(max childen of each node), m > 2
// l is the cap of leaf, l > 2
func New(m int, l int, cmp Comparer) *BplusTree {
	tree := &BplusTree{}
	tree.m = m
	tree.l = l
	tree.cmp = cmp
	if tree.m < 3 {
		tree.m = 3
	}
	if tree.l < 3 {
		tree.l = 3
	}
	return tree
}

func (tree *BplusTree) Insert(key, val interface{}) {
	if tree.root == nil {
		l := tree.newLeaf()
		l.append(key, val)
		tree.root = l
		tree.size++
		return
	}

	l, pos := search(tree.root, key, tree.cmp)
	if pos >= 0 {
		l.items[pos].val = val
		return
	}

	tree.size++
	l.insert(key, val, tree.cmp)
	if l.full() {
		tree.splitLeaf(l)
	}
}

func (tree *BplusTree) Del(key interface{}) {
	l, pos := search(tree.root, key, tree.cmp)
	if pos < 0 {
		return
	}

	tree.size--
	l.del(pos)
	if !l.underflow() {
		return
	}

	var left, right *leaf
	var maxCount int
	if l.next != nil && l.parent == l.next.parent {
		left = l
		right = l.next
		maxCount = right.count
	} else {
		left = l.pre
		right = l
		maxCount = left.count
	}

	parent := left.parent
	var leftChildPos int
	for i := 0; i < parent.childCount; i++ {
		if parent.children[i] == left {
			leftChildPos = i
			break
		}
	}

	if maxCount > tree.l/2 { //  more than half-full
		tree.redistributeLeaf(left, right, leftChildPos)
	} else {
		tree.mergeLeaf(left, right, leftChildPos)
	}
}

func (tree *BplusTree) splitLeaf(l *leaf) {
	newLeaf := tree.newLeaf()
	middleKey := l.split(newLeaf)

	parent := l.parent
	if parent == nil {
		tree.newRoot(middleKey, l, newLeaf)
		return
	}

	pos := parent.insertKey(middleKey, tree.cmp)
	if pos == parent.keyCount-1 {
		parent.children[parent.childCount] = newLeaf
	} else {
		copy(parent.children[pos+1:], parent.children[pos:parent.childCount])
		parent.children[pos] = l
		parent.children[pos+1] = newLeaf
	}
	parent.childCount++

	if parent.full() {
		tree.splitNode(parent)
	}
}

func (tree *BplusTree) splitNode(a *node) {
	b := tree.newNode()
	middleKey := a.split(b)

	if a.parent != nil {
		pos := a.parent.insertKey(middleKey, tree.cmp)
		if pos == a.parent.keyCount-1 {
			a.parent.children[a.parent.childCount] = b
		} else {
			size := a.parent.childCount
			copy(a.parent.children[pos+1:], a.parent.children[pos:size])
			a.parent.children[pos] = a
			a.parent.children[pos+1] = b
		}
		a.parent.childCount++

		b.parent = a.parent
	} else {
		tree.newRoot(middleKey, a, b)
	}

	if a.parent.full() {
		tree.splitNode(a.parent)
	}
}

func (tree *BplusTree) redistributeLeaf(left *leaf, right *leaf, leftChildPos int) {
	if left.count < right.count {
		i := 0
		for left.underflow() {
			left.append(right.items[i].key, right.items[i].val)
			i++
		}

		copy(right.items[:], right.items[i:right.count])
		right.count = right.count - i
	} else {
		diff := tree.l/2 - right.count
		copy(right.items[diff:], right.items[:right.count])
		j := left.count - 1
		for i := 0; i < diff; i++ {
			right.items[i] = left.items[j]
			j--
		}
		left.count -= diff
		right.count += diff
	}

	left.parent.keys[leftChildPos] = right.items[0].key
}

func (tree *BplusTree) mergeLeaf(left *leaf, right *leaf, leftChildPos int) {
	for i := 0; i < right.count; i++ {
		left.append(right.items[i].key, right.items[i].val)
	}

	parent := left.parent
	copy(parent.keys[leftChildPos:], parent.keys[leftChildPos+1:parent.keyCount])
	copy(parent.children[leftChildPos+1:], parent.children[leftChildPos+2:parent.childCount])
	parent.keyCount--
	parent.childCount--

	left.next = right.next
	right.next.pre = left

	if parent.underflow() {

	}
}

func (tree *BplusTree) newRoot(key interface{}, left, right commonNode) {
	a := tree.newNode()
	a.appendKey(key)
	left.setParent(a)
	right.setParent(a)
	a.appendChild(left)
	a.appendChild(right)

	tree.root = a
}

func search(from commonNode, key interface{}, cmp Comparer) (*leaf, int) {

	for {
		if from.isLeaf() {
			break
		}

		a := from.(*node)
		pos, found := a.find(key, cmp)
		if found || cmp(key, a.keys[pos]) > 0 {
			pos++
		}
		from = a.children[pos]
	}

	l := from.(*leaf)
	pos, found := l.find(key, cmp)
	if !found {
		return l, -1
	}

	return l, pos
}
