package bst

type RBTree struct {
	UnbalancedTree
}

const (
	rbRed   = 1
	rbBlack = 2
)

func NewRB(comparer Comparer) RBTree {
	tree := RBTree{}
	tree.cmp = comparer

	return tree
}

func (tree *RBTree) Insert(k, v interface{}) {
	x, exist := insert(tree.root, k, v, tree.cmp)
	if !exist {
		tree.size++
	}

	if tree.root == nil {
		tree.root = x
		return
	}

	if exist {
		return
	}

	x.setColor(rbRed)
	tree.insertFixup(x)
}

func (tree *RBTree) insertFixup(x *node) {
	for x != tree.root && x.parent.color() == rbRed {
		if x.parent == x.parent.parent.left {
			// y is x's right uncle
			y := x.parent.parent.right
			if y != nil && y.color() == rbRed {
				// case 1 - change the colors
				x.parent.setColor(rbBlack)
				y.setColor(rbBlack)
				x.parent.parent.setColor(rbRed)
				x = x.parent.parent
			} else {

				if x == x.parent.right {
					// uncle is black and x is to the right
					// case 2
					rotateNode := x.parent
					leftRotate(rotateNode, &(tree.root), false)
					x = rotateNode
				}

				// case 3
				x.parent.setColor(rbBlack)
				x.parent.parent.setColor(rbRed)
				rotateNode := x.parent.parent
				rightRotate(rotateNode, &(tree.root), false)
			}
		} else {
			y := x.parent.parent.left
			if y != nil && y.color() == rbRed {
				x.parent.setColor(rbBlack)
				y.setColor(rbBlack)
				x.parent.parent.setColor(rbRed)
				x = x.parent.parent
			} else {

				if x == x.parent.left {
					rotateNode := x.parent
					rightRotate(rotateNode, &(tree.root), false)
					x = rotateNode
				}

				x.parent.setColor(rbBlack)
				x.parent.parent.setColor(rbRed)
				rotateNode := x.parent.parent
				leftRotate(rotateNode, &(tree.root), false)
			}
		}

	}

	tree.root.setColor(rbBlack)
}

func (tree *RBTree) Del(key interface{}) bool {
	px, x, y, leftChild := del(&(tree.root), key, tree.cmp)
	if y == nil {
		return false
	}

	tree.size--

	if px != nil && y.color() == rbBlack {
		tree.delFixup(px, x, leftChild) // x maybe nil
	}

	return true
}

func (tree *RBTree) delFixup(px *node, x *node, leftChild bool) {
	for x != tree.root && isBlack(x) {
		if leftChild {
			// case 1
			w := px.right // w is x's sibling

			if isRed(w) {
				w.setColor(rbBlack)
				px.setColor(rbRed)
				leftRotate(px, &(tree.root), false)
				w = px.right
			}

			if w != nil && isBlack(w.left) && isBlack(w.right) {
				// case 2
				w.setColor(rbRed)
				x = px
				px = x.parent
				if px != nil && x == px.left {
					leftChild = true
				} else {
					leftChild = false
				}
			} else {
				if w != nil && isBlack(w.right) {
					// case 3
					if w.left != nil {
						w.left.setColor(rbBlack)
					}
					rightRotate(w, &(tree.root), false)
					w.setColor(rbRed)
					w = px.right
				}

				// case 4
				if w != nil {
					w.setColor(px.color())
					if w.right != nil {
						w.right.setColor(rbBlack)

					}
				}
				px.setColor(rbBlack)
				if px.right != nil {
					leftRotate(px, &(tree.root), false)
				}

				x = tree.root
			}
		} else {
			w := px.left

			if isRed(w) {
				w.setColor(rbBlack)
				px.setColor(rbRed)
				rightRotate(px, &(tree.root), false)
				w = px.left
			}

			if w != nil && isBlack(w.left) && isBlack(w.right) {
				w.setColor(rbRed)
				x = px
				px = x.parent
				if px != nil && x == px.left {
					leftChild = true
				} else {
					leftChild = false
				}
			} else {
				if w != nil && isBlack(w.left) {
					if w.right != nil {
						w.right.setColor(rbBlack)
					}
					leftRotate(w, &(tree.root), false)
					w.setColor(rbRed)
					w = px.left
				}

				if w != nil {
					w.setColor(px.color())
					if w.left != nil {
						w.left.setColor(rbBlack)
					}
				}
				if px.left != nil {
					rightRotate(px, &(tree.root), false)
				}
				px.setColor(rbBlack)

				x = tree.root
			}

		}
	}

	if x != nil {
		x.setColor(rbBlack)
	}
}

func isRed(x *node) bool {
	if x != nil && x.color() == rbRed {
		return true
	}

	return false
}

func isBlack(x *node) bool {
	if x == nil {
		return true
	}

	if x.color() == rbBlack {
		return true
	}

	return false
}
