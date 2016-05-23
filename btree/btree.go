package btree

type Comparer func(interface{}, interface{}) int

type BTree struct {
	cmp  Comparer
	m    int
	size int
	root *node
}

// m is order(max childen of each node), m > 2
func New(m int, cmp Comparer) *BTree {
	tree := &BTree{}
	tree.m = m
	tree.cmp = cmp
	if tree.m < 3 {
		tree.m = 3
	}
	return tree
}

func (tree *BTree) Size() int {
	return tree.size
}

func (tree *BTree) Clear() {
	tree.root = nil
	tree.size = 0
}

func (tree *BTree) Find(key interface{}) (val interface{}, found bool) {
	a, pos := search(tree.root, key, tree.cmp)
	if pos < 0 {
		return
	}

	found = true
	val = a.items[pos].val
	return
}

func (tree *BTree) Insert(key, val interface{}) {
	if tree.root == nil {
		tree.root = tree.newNode()
		tree.root.appendItem(key, val)
		tree.size++
		return
	}

	a, pos := search(tree.root, key, tree.cmp)
	if pos >= 0 {
		a.items[pos].val = val
		return
	}

	tree.size++

	a.insertItem(key, val, tree.cmp)
	if a.isFull() {
		tree.split(a)
	}
}

func (tree *BTree) Del(key interface{}) {
	a, pos := search(tree.root, key, tree.cmp)
	if pos < 0 {
		return
	}

	tree.size--

	// If the value to be deleted does not occur in a leaf, we replace it with
	// the largest value in its left subtree and then proceed to delete that
	// value from the node that originally contained it.
	if a.children != nil {
		largest := findLargest(a.children[pos])
		a.items[pos] = largest.items[largest.itemCount-1]
		a = largest
		pos = largest.itemCount - 1
	}

	a.delItem(pos)

	// empty tree
	if a == tree.root && a.itemLen() == 0 && a.childCount == 0 {
		tree.root = nil
		return
	}

	if tree.needRebalance(a) {
		tree.rebalance(a)
	}
}

func (tree *BTree) needRebalance(a *node) bool {
	// we require the root to have at least 1 value in it and
	// all other nodes to have at least (M-1)/2 values in them.
	if a.itemLen() >= (tree.m-1)/2 {
		return false
	}
	if a == tree.root && a.itemLen() > 0 {
		return false
	}

	return true
}

// a.len() should be tree.m
func (tree *BTree) split(a *node) {
	b := tree.newNode()

	m := a.itemLen()

	// copy the last (M-1)/2 values
	for i := (m-1)/2 + 1; i < m; i++ {
		b.appendItem(a.items[i].key, a.items[i].val)
	}

	middle := a.items[(m-1)/2]

	a.truncItem((m - 1) / 2)

	if a.children != nil {
		count := a.childCount
		a.childCount = a.itemLen() + 1

		b.children = make([]*node, tree.m+1)
		for i := a.childCount; i < count; i++ {
			b.appendChild(a.children[i])
			a.children[i].parent = b
		}
	}

	if a.parent != nil {
		pos := a.parent.insertItem(middle.key, middle.val, tree.cmp)
		if pos == a.parent.itemLen()-1 {
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
		tree.root = tree.newNode()
		tree.root.children = make([]*node, tree.m+1)
		tree.root.appendItem(middle.key, middle.val)
		tree.root.children[0] = a
		tree.root.children[1] = b
		tree.root.childCount = 2

		a.parent = tree.root
		b.parent = tree.root
	}

	if a.parent.isFull() {
		tree.split(a.parent)
	}
}

func (tree *BTree) rebalance(a *node) {
	parent := a.parent

	var childPos int
	for childPos = 0; childPos < parent.childCount; childPos++ {
		if parent.children[childPos] == a {
			break
		}
	}

	var left, right *node
	var itemPos int
	if childPos == parent.childCount-1 { // the last child
		left = parent.children[childPos-1]
		right = a
		itemPos = childPos - 1
	} else {
		left = a
		right = parent.children[childPos+1]
		itemPos = childPos
	}

	if left.itemLen()+right.itemLen()+1 >= tree.m {
		redistribute(left, right, parent, itemPos)
	} else {
		tree.merge(left, right, parent, itemPos)
	}
}

func redistribute(left, right, parent *node, itemPos int) {
	items := make([]item, left.itemLen()+right.itemLen()+1)

	offset := 0
	copy(items[offset:], left.items[:left.itemLen()])
	offset += left.itemLen()
	items[offset] = parent.items[itemPos]
	offset += 1
	copy(items[offset:], right.items[:right.itemLen()])

	offset = 0
	size := (left.itemLen() + right.itemLen()) / 2
	copy(left.items[:], items[offset:size])
	offset += size
	parent.items[itemPos] = items[offset]
	offset += 1
	copy(right.items[:], items[offset:])

	left.itemCount = size
	right.itemCount = len(items) - size - 1

	if left.childCount+right.childCount > 0 {
		if left.childCount < right.childCount {
			i := 0
			for left.childCount < left.itemCount+1 {
				left.appendChild(right.children[i])
				right.children[i].parent = left
				i++
			}
			copy(right.children[:], right.children[i:right.childCount])
			right.childCount = right.childCount - i
		} else {
			diff := right.itemCount + 1 - right.childCount
			copy(right.children[diff:], right.children[:right.childCount])
			j := left.childCount - 1
			for i := 0; i < diff; i++ {
				right.children[i] = left.children[j]
				right.children[i].parent = right
				j--
			}

			left.childCount = j + 1
			right.childCount = right.itemCount + 1
		}
	}
}

func (tree *BTree) merge(left, right, parent *node, itemPos int) {
	left.appendItem(parent.items[itemPos].key, parent.items[itemPos].val)

	for i := 0; i < right.itemLen(); i++ {
		left.appendItem(right.items[i].key, right.items[i].val)
	}

	for i := 0; i < right.childCount; i++ {
		left.appendChild(right.children[i])
		right.children[i].parent = left
	}

	// root underflow
	if parent == tree.root && parent.childCount == 2 {
		tree.root = left
		tree.root.parent = nil
	} else {
		copy(parent.items[itemPos:], parent.items[itemPos+1:])
		copy(parent.children[itemPos:], parent.children[itemPos+1:])
		parent.children[itemPos] = left
		parent.itemCount--
		parent.childCount--

		if tree.needRebalance(parent) {
			tree.rebalance(parent)
		}
	}
}

func (tree *BTree) newNode() *node {
	a := &node{}
	a.items = make([]item, tree.m)
	return a
}
