package bplustree

type commonNode interface {
	setParent(a *node)
	getParent() *node
	isLeaf() bool
	getKey(i int) interface{}
	size() int
}
