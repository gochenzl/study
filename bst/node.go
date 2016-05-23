package bst

type Comparer func(interface{}, interface{}) int

type node struct {
	left     *node
	right    *node
	parent   *node
	key      interface{}
	val      interface{}
	addition int // height for avl, color for red-black
}

func (x *node) height() int {
	if x == nil {
		return 0
	}

	return x.addition
}

func (x *node) setHeight(h int) {
	x.addition = h
}

func (x *node) color() int {
	return x.addition
}

func (x *node) setColor(color int) {
	x.addition = color
}

// return the node inserted
func insert(x *node, k, v interface{}, cmp Comparer) (retNode *node, exist bool) {
	if x == nil {
		retNode = &node{key: k, val: v}
		return
	}

	cur := x
	parent := x.parent

	for {
		ret := cmp(k, cur.key)
		if ret == 0 {
			cur.val = v
			retNode = cur
			exist = true
			return
		}

		parent = cur
		if ret < 0 {
			cur = cur.left
		} else {
			cur = cur.right
		}

		if cur == nil {
			retNode = &node{parent: parent, key: k, val: v}
			if ret < 0 {
				parent.left = retNode
			} else {
				parent.right = retNode
			}

			return
		}
	}

}

func delWithOneChild(proot **node, x *node, child *node) {
	parent := x.parent
	if parent != nil {
		if parent.left == x {
			parent.left = child
		} else {
			parent.right = child
		}
		child.parent = parent
	} else {
		*proot = child
	}
}

func delWithNoChildren(proot **node, x *node) {
	parent := x.parent
	if parent != nil {
		if parent.left == x {
			parent.left = nil
		} else {
			parent.right = nil
		}
	} else { // delete the root
		*proot = nil
	}
}

// 1.Deleting a node with no children: simply remove the node from the tree.
//
// 2.Deleting a node with one child: remove the node and replace it with its child.
//
// 3.Deleting a node with two children: call the node to be deleted N. Do not delete N.
// Instead, choose either its in-order successor node or its in-order predecessor node, R.
// Copy the value of R to N, then delete R
//
// y is node deleted
// x is child of y
// parent is parent of y, after delete is parent of x
func del(proot **node, key interface{}, cmp Comparer) (parent *node, x *node, y *node, leftChild bool) {
	y = search(*proot, key, cmp)
	if y == nil {
		return
	}

	if y.left != nil && y.right != nil {
		successor := getSuccessor(y)
		y.key = successor.key
		y.val = successor.val
		y = successor
	}

	//     parent
	//       /
	//      y
	//     /
	//    x

	if y.left != nil {
		x = y.left
	} else {
		x = y.right
	}

	if x != nil {
		x.parent = y.parent
	}

	parent = y.parent
	if parent != nil {
		if parent.left == y {
			leftChild = true
			parent.left = x
		} else {
			parent.right = x
		}
	} else {
		*proot = x
	}

	return
}

func search(root *node, key interface{}, cmp Comparer) *node {
	if root == nil {
		return nil
	}

	x := root
	for x != nil {
		ret := cmp(key, x.key)
		if ret == 0 {
			return x
		} else if ret < 0 {
			x = x.left
		} else {
			x = x.right
		}
	}

	return nil
}

// find left most
func findMin(x *node) *node {
	for {
		if x.left == nil {
			break
		}

		x = x.left
	}

	return x
}

func getSuccessor(x *node) *node {
	if x.right != nil {
		return findMin(x.right)
	}

	y := x.parent
	for y != nil && x == y.right {
		x = y
		y = y.parent
	}

	return y
}

//Rotates node a to the left, making its right child into its parent.
//
//     A              B
//    / \            / \
//   X   B  -->     A   Z
//       /\        / \   \
//      Y  Z      X   Y   W
//          \
//          W
//
func leftRotate(a *node, proot **node, h bool) {
	parent := a.parent
	b := a.right

	a.right = b.left
	if b.left != nil {
		b.left.parent = a
	}

	b.left = a
	b.parent = parent
	a.parent = b

	if h {
		a.setHeight(calcHeight(a))
		b.setHeight(calcHeight(b))
	}

	// attach b to parent
	if parent != nil {
		if parent.left == a {
			parent.left = b
		} else {
			parent.right = b
		}

		if h {
			updateHeight(parent)
		}
	} else {
		*proot = b
	}
}

//Rotates node a to the right, making its left child into its parent.
//
//        A           B
//       / \         / \
//      B   Z  -->  X   A
//     / \         /   / \
//    X   Y       W   Y   Z
//   /
//  W
//
func rightRotate(a *node, proot **node, h bool) {
	parent := a.parent
	b := a.left

	a.left = b.right
	if b.right != nil {
		b.right.parent = a
	}

	b.right = a
	b.parent = parent
	a.parent = b

	if h {
		a.setHeight(calcHeight(a))
		b.setHeight(calcHeight(b))
	}

	// attach b to parent
	if parent != nil {
		if parent.left == a {
			parent.left = b
		} else {
			parent.right = b
		}

		if h {
			updateHeight(parent)
		}
	} else {
		*proot = b
	}
}

func updateHeight(x *node) {
	for x != nil {
		x.setHeight(calcHeight(x))
		x = x.parent
	}
}

func calcHeight(x *node) int {
	leftHeight := x.left.height()
	rightHeight := x.right.height()

	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}
