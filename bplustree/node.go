package bplustree

type node struct {
	parent     *node
	keys       []interface{}
	keyCount   int
	children   []commonNode
	childCount int
}

func (a *node) setParent(parent *node) {
	a.parent = parent
}

func (a *node) getParent() *node {
	return a.parent
}

func (a *node) isLeaf() bool {
	return false
}

func (a *node) getKey(i int) interface{} {
	return a.keys[i]
}

func (a *node) size() int {
	return a.keyCount
}

// the same of btree
func (a *node) split(b *node) (middleKey interface{}) {
	m := a.keyCount

	// copy the last (M-1)/2 values
	for i := (m-1)/2 + 1; i < m; i++ {
		b.appendKey(a.keys[i])
	}

	middleKey = a.keys[(m-1)/2]

	a.truncKey((m - 1) / 2)

	childCount := a.childCount
	a.childCount = a.keyCount + 1

	for i := a.childCount; i < childCount; i++ {
		b.appendChild(a.children[i])
		a.children[i].setParent(b)
	}

	return
}

func (a *node) full() bool {
	return a.keyCount == len(a.keys)
}

func (a *node) underflow() bool {
	return a.keyCount < len(a.keys)/2
}

func (a *node) appendKey(key interface{}) {
	a.keys[a.keyCount] = key
	a.keyCount++
}

func (a *node) appendChild(child commonNode) {
	a.children[a.childCount] = child
	a.childCount++
}

func (a *node) truncKey(newLen int) {
	a.keyCount = newLen
}

func (a *node) insertKey(key interface{}, cmp Comparer) int {
	pos, found := a.find(key, cmp)
	if found {
		return pos
	}

	if cmp(key, a.keys[pos]) > 0 {
		pos++
	}

	a.appendKey(key)
	if pos < a.keyCount {
		copy(a.keys[pos+1:a.keyCount], a.keys[pos:a.keyCount-1])
		a.keys[pos] = key
	}

	return pos
}

func (a *node) find(key interface{}, cmp Comparer) (i int, found bool) {
	l := 0
	h := a.keyCount - 1
	for l <= h {
		i = (l + h) >> 1
		ret := cmp(key, a.keys[i])
		if ret < 0 {
			h = i - 1
		} else if ret > 0 {
			l = i + 1
		} else {
			found = true
			return
		}
	}

	return
}

func (tree *BplusTree) newNode() *node {
	a := &node{}
	a.keys = make([]interface{}, tree.m)
	a.children = make([]commonNode, tree.m+1)
	return a
}
