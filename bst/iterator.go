package bst

import "github.com/gochenzl/study/stack"

type Iterator struct {
	st  *stack.Stack
	cur *node
}

func (tree *UnbalancedTree) NewIterator() Iterator {
	var iter Iterator
	if tree.root == nil {
		return iter
	}

	iter.st = stack.New(100)
	cur := tree.root
	for cur != nil {
		iter.st.Push(cur)
		cur = cur.left
	}

	return iter
}

func (iter *Iterator) Next() bool {
	if iter.st == nil {
		return false
	}

	val, ok := iter.st.Pop()
	if !ok {
		return false
	}

	iter.cur = val.(*node)
	n := iter.cur.right
	for n != nil {
		iter.st.Push(n)
		n = n.left
	}

	return true
}

func (iter *Iterator) Key() interface{} {
	if iter.cur != nil {
		return iter.cur.key
	}

	return nil
}

func (iter *Iterator) Value() interface{} {
	if iter.cur != nil {
		return iter.cur.val
	}

	return nil
}
