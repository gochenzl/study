package bplustree

type item struct {
	key interface{}
	val interface{}
}

type leaf struct {
	items  []item
	count  int
	parent *node
	pre    *leaf
	next   *leaf
}

func (a *leaf) setParent(parent *node) {
	a.parent = parent
}

func (a *leaf) getParent() *node {
	return a.parent
}

func (a *leaf) isLeaf() bool {
	return true
}

func (a *leaf) getKey(i int) interface{} {
	return a.items[i].key
}

func (a *leaf) size() int {
	return a.count
}

func (a *leaf) split(b *leaf) (middleKey interface{}) {
	size := a.count

	middleKey = a.items[size/2].key
	for i := size / 2; i < size; i++ {
		b.append(a.items[i].key, a.items[i].val)
	}
	a.trunc(size / 2)

	b.next = a.next
	if b.next != nil {
		b.next.pre = b
	}
	b.pre = a
	b.parent = a.parent
	a.next = b

	return
}

func (a *leaf) find(key interface{}, cmp Comparer) (i int, found bool) {
	l := 0
	h := a.count - 1
	for l <= h {
		i = (l + h) >> 1
		ret := cmp(key, a.items[i].key)
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

func (a *leaf) append(key interface{}, val interface{}) {
	a.items[a.count] = item{key, val}
	a.count++
}

func (a *leaf) insert(key interface{}, val interface{}, cmp Comparer) {
	pos, found := a.find(key, cmp)
	if found {
		return
	}

	if cmp(key, a.items[pos].key) > 0 {
		pos++
	}

	a.append(key, val)
	if pos < a.count {
		copy(a.items[pos+1:a.count], a.items[pos:a.count-1])
		a.items[pos] = item{key, val}
	}
}

func (a *leaf) del(pos int) {
	if pos == a.count-1 {
		a.count--
		return
	}

	copy(a.items[pos:a.count], a.items[pos+1:])
	a.count--
}

func (a *leaf) trunc(newLen int) {
	a.count = newLen
}

func (a *leaf) full() bool {
	if a.count == len(a.items) {
		return true
	}

	return false
}

func (a *leaf) underflow() bool {
	return a.count < len(a.items)/2
}

func (tree *BplusTree) newLeaf() *leaf {
	l := &leaf{}
	l.items = make([]item, tree.l)
	return l
}
