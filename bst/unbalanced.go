package bst

type UnbalancedTree struct {
	cmp  Comparer
	root *node
	size int
}

func NewUnbalanced(comparer Comparer) UnbalancedTree {
	tree := UnbalancedTree{}
	tree.cmp = comparer
	return tree
}

func (tree *UnbalancedTree) Insert(k, v interface{}) {
	n, exist := insert(tree.root, k, v, tree.cmp)
	if !exist {
		tree.size++
	}
	if tree.root == nil {
		tree.root = n
	}
}

func (tree *UnbalancedTree) Del(key interface{}) bool {
	_, _, y, _ := del(&(tree.root), key, tree.cmp)
	if y != nil {
		tree.size--
		return true
	}

	return false
}

func (tree *UnbalancedTree) Clear() {
	tree.root = nil
	tree.size = 0
}

func (tree *UnbalancedTree) Find(key interface{}) (val interface{}, found bool) {
	n := search(tree.root, key, tree.cmp)
	if n == nil {
		return
	}

	val = n.val
	found = true
	return
}

func (tree *UnbalancedTree) Size() int {
	return tree.size
}
