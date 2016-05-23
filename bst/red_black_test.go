package bst

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRBTree(t *testing.T) {
	size := 100000
	keys := make([]int, 0, size)

	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < size; i++ {
		k := r.Int() % (size * 2)
		keys = append(keys, k)
	}

	tree := NewRB(intCompare)
	for i := 0; i < size; i++ {
		tree.Insert(keys[i], 1)
	}

	if !checkRBTree(tree.root) {
		t.Error("checkRBTree")
	}

	for i := 0; i < size/2; i++ {
		key := keys[r.Int()%size]
		tree.Del(key)
	}

	if !checkRBTree(tree.root) {
		t.Error("checkRBTree")
	}

}

//A red-black tree is a balanced binary search tree with the following properties:
//1.Every node is colored red or black.
//2.Every leaf is a NIL node, and is colored black.
//3.If a node is red, then both its children are black.
//4.Every simple path from a node to a descendant leaf contains the same number of black nodes.
func checkRBTree(n *node) bool {
	if n == nil {
		return true
	}

	if n.color() == rbRed {
		if n.left != nil && n.left.color() != rbBlack {
			return false
		}

		if n.right != nil && n.right.color() != rbBlack {
			return false
		}
	}

	if !checkBlackHeight(n) {
		return false
	}

	if !checkRBTree(n.left) {
		return false
	}

	if !checkRBTree(n.right) {
		return false
	}

	return true
}

func checkBlackHeight(n *node) bool {
	path := make([]*node, 1000)
	index := 0
	heights := make([]int, 0, 100)
	calcBlackHeight(n, path, index, &heights)

	for i := 0; i < len(heights)-1; i++ {
		if heights[i] != heights[i+1] {
			return false
		}
	}

	return true
}

func calcBlackHeight(n *node, path []*node, index int, heights *[]int) {
	if n == nil {
		return
	}

	path[index] = n
	index++

	if n.left == nil && n.right == nil {
		var h int
		for i := 0; i < index; i++ {
			if path[i].color() == rbBlack {
				h++
			}
		}
		*heights = append(*heights, h)
		return
	} else {
		calcBlackHeight(n.left, path, index, heights)
		calcBlackHeight(n.right, path, index, heights)
	}
}

func rbString(k, v interface{}, addition int) string {
	a := k.(int)
	if addition == rbRed {
		return strconv.Itoa(a) + "(r)"
	} else {
		return strconv.Itoa(a) + "(b)"
	}

	panic("hello")
}
