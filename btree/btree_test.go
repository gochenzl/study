package btree

import (
	"container/list"
	"math/rand"
	"testing"
	"time"
)

func TestBTreeInsert(t *testing.T) {
	tree := New(3, intCompare)

	//                -----
	//                | 4 |
	//                -----
	//              /       \
	//         -----         -----
	//         | 2 |         | 6 |
	//         -----         -----
	//        /     \       /     \
	//   -----   -----  -----     -----
	//   | 1 |   | 3 |  | 5 |     | 7 |
	//   -----   -----  -----     -----
	tree.Insert(1, 1)
	tree.Insert(2, 1)
	tree.Insert(3, 1)
	tree.Insert(4, 1)
	tree.Insert(5, 1)
	tree.Insert(6, 1)
	tree.Insert(7, 1)

	if !checkNodeItems(nil, tree.root, []item{item{4, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{2, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{6, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[0], []item{item{1, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[1], []item{item{3, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[0], []item{item{5, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[1], []item{item{7, 1}}) {
		t.Error("Insert")
	}

	if tree.Size() != 7 {
		t.Error("Size")
	}

	tree.Insert(7, 10)
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[1], []item{item{7, 10}}) {
		t.Error("Insert")
	}

	tree = New(5, intCompare)
	tree.Insert(5, 1)
	tree.Insert(7, 1)
	tree.Insert(10, 1)
	tree.Insert(22, 1)
	tree.Insert(44, 1)
	tree.Insert(2, 1)
	tree.Insert(3, 1)
	tree.Insert(50, 1)
	tree.Insert(55, 1)
	tree.Insert(66, 1)
	tree.Insert(45, 1)
	tree.Insert(68, 1)
	tree.Insert(70, 1)

	tree.Insert(17, 1)
	tree.Insert(6, 1)
	tree.Insert(21, 1)
	tree.Insert(67, 1)

	if !checkNodeItems(nil, tree.root, []item{item{22, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{5, 1}, item{10, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{50, 1}, item{67, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[0], []item{item{2, 1}, item{3, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[1], []item{item{6, 1}, item{7, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[2], []item{item{17, 1}, item{21, 1}}) {
		t.Error("Insert")
	}

	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[0], []item{item{44, 1}, item{45, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[1], []item{item{55, 1}, item{66, 1}}) {
		t.Error("Insert")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[2], []item{item{68, 1}, item{70, 1}}) {
		t.Error("Insert")
	}

	if tree.Size() != 17 {
		t.Error("Size")
	}
}

func TestBTreeDel(t *testing.T) {
	//                    -----------------------------------
	//                    |   5     |    11     |    17     |
	//					  -----------------------------------
	//				     /          /             \           \
	//	        ----------     ---------       -----------     -----------
	//	        |  1 | 3 |     | 7 | 9 |       | 13 | 15 |     | 19 | 21 |
	//	        ----------     ---------       -----------     -----------
	keys := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}
	tree := New(5, intCompare)
	createTree(tree, keys)

	//                    -------------------------
	//                    |    11     |    17     |
	//					  -------------------------
	//				     /            \            \
	//	        -------------------   -----------   -----------
	//	        |  1 | 5 | 7 | 9 |    | 13 | 15 |   | 19 | 21 |
	//	        -------------------   -----------   -----------
	tree.Del(3)
	if !checkNodeItems(nil, tree.root, []item{item{11, 1}, item{17, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{1, 1}, item{5, 1}, item{7, 1}, item{9, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{13, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[2], []item{item{19, 1}, item{21, 1}}) {
		t.Error("Del")
	}

	if tree.Size() != 10 {
		t.Error("Size")
	}

	//                    -------------
	//                    |    11     |
	//					  -------------
	//				     /            \
	//	        ------------------   ---------------------
	//	        |  1 | 5 | 7 | 9 |   | 15 | 17 | 19 | 21 |
	//	        -------------------  ---------------------
	tree.Del(13)
	if !checkNodeItems(nil, tree.root, []item{item{11, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{1, 1}, item{5, 1}, item{7, 1}, item{9, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{15, 1}, item{17, 1}, item{19, 1}, item{21, 1}}) {
		t.Error("Del")
	}
	if tree.Size() != 9 {
		t.Error("Size")
	}

	//                    -------------
	//                    |     7     |
	//					  -------------
	//				     /            \
	//	        ----------      ---------------
	//	        |  1 | 5 |      | 9 | 11 | 15 |
	//	        ----------      ---------------
	tree.Del(17)
	tree.Del(19)
	tree.Del(21)
	if !checkNodeItems(nil, tree.root, []item{item{7, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{1, 1}, item{5, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{9, 1}, item{11, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if tree.Size() != 6 {
		t.Error("Size")
	}

	//                    -------------
	//                    |     9     |
	//					  -------------
	//				     /            \
	//	        ----------       -----------
	//	        |  1 | 5 |       | 11 | 15 |
	//	        ----------       -----------
	tree.Del(7)
	if !checkNodeItems(nil, tree.root, []item{item{9, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{1, 1}, item{5, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{11, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if tree.Size() != 5 {
		t.Error("Size")
	}

	//	        ---------- --------
	//	        |  1 | 5 |11 | 15 |
	//	        -------------------
	tree.Del(9)
	if !checkNodeItems(nil, tree.root, []item{item{1, 1}, item{5, 1}, item{11, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if tree.Size() != 4 {
		t.Error("Size")
	}

	tree.Del(1)
	tree.Del(11)
	tree.Del(15)
	tree.Del(5)
	if tree.root != nil {
		t.Error("Del")
	}
	if tree.Size() != 0 {
		t.Error("Size")
	}

	//                     -----------------------------------
	//                     |                11               |
	//					   -----------------------------------
	//					  /                                   \
	//			-----------------                              -----------------------------
	//			|   4   |   7   |                              |    16   |   20   |   30   |
	//			-----------------                              -----------------------------
	//		   /        |        \                            /         /          \        \
	// ---------      ---------    ----------       -----------  -----------    -----------   -----------
	// | 1 | 3 |      | 5 | 6 |    | 8 | 10 |       | 13 | 14 |  | 17 | 18 |    | 24 | 25 |   | 31 | 35 |
	// --------- 	   ---------   ----------       -----------  -----------    -----------   -----------
	keys = []int{1, 3, 4, 5, 6, 7, 8, 10, 11, 13, 14, 16, 17, 18, 20, 24, 25, 30, 31, 35}
	tree.Clear()
	createTree(tree, keys)

	//                     -----------------------------------
	//                     |                16               |
	//					   -----------------------------------
	//					  /                                   \
	//			----------------------------                   -------------------
	//			|      7     |      11     |                   |   20   |   30   |
	//			----------------------------                   -------------------
	//		   /             |              \                 /         |         \
	// -----------------  ----------     -----------    -----------  -----------   -----------
	// | 3 | 4 | 5 | 6 |  | 8 | 10 |     | 13 | 14 |    | 17 | 18 |  | 24 | 25 |   | 31 | 35 |
	// -----------------  ----------     -----------    -----------  -----------   -----------
	tree.Del(1)
	if !checkNodeItems(nil, tree.root, []item{item{16, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{7, 1}, item{11, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{20, 1}, item{30, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[0], []item{item{3, 1}, item{4, 1}, item{5, 1}, item{6, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[1], []item{item{8, 1}, item{10, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[2], []item{item{13, 1}, item{14, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[0], []item{item{17, 1}, item{18, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[1], []item{item{24, 1}, item{25, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[2], []item{item{31, 1}, item{35, 1}}) {
		t.Error("Del")
	}

	//			--------------------------------------------------------
	//			|      7     |      11     |       16     |      20    |
	//			--------------------------------------------------------
	//		   /             |              \              \            \
	// -----------------  ----------     -----------    -----------      ---------------------
	// | 3 | 4 | 5 | 6 |  | 8 | 10 |     | 13 | 14 |    | 17 | 18 |      | 24 | 25 | 30 | 31 |
	// -----------------  ----------     -----------    -----------      ---------------------
	tree.Del(35)
	if !checkNodeItems(nil, tree.root, []item{item{7, 1}, item{11, 1}, item{16, 1}, item{20, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{3, 1}, item{4, 1}, item{5, 1}, item{6, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{8, 1}, item{10, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[2], []item{item{13, 1}, item{14, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[3], []item{item{17, 1}, item{18, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[4], []item{item{24, 1}, item{25, 1}, item{30, 1}, item{31, 1}}) {
		t.Error("Del")
	}

	//                               -----------------------------------
	//                               |                16               |
	//					             -----------------------------------
	//					            /                                   \
	//   	--------------------------------------                       ---------------------
	//   	|     3     |     9     |     12     |                       |    21   |   25    |
	//	    --------------------------------------                       ---------------------
	//	   /           /             \            \                      /         |         \
	// --------  ---------       -----------   ----------       -----------  -----------    -----------
	// | 1 | 2|  | 6 | 8 |       | 10 | 11 |   | 13 | 15 |      | 18 | 19 |  | 22 | 23 |    | 29 | 30 |
	// --------  ---------       -----------   ----------       -----------  -----------    -----------
	keys = []int{6, 8, 9, 10, 11, 12, 13, 15, 16, 18, 19, 21, 22, 23, 25, 29, 30, 1, 2, 3}
	tree.Clear()
	createTree(tree, keys)

	//                               -----------------------------------
	//                               |                12               |
	//					             -----------------------------------
	//					            /                                   \
	//   	------------------------                          --------------------
	//   	|     3     |     9     |                         |    16   |   21    |
	//	    -------------------------                         ---------------------
	//	   /           /             \                       /         |           \
	// --------  ---------       -----------        ----------    -----------       ---------------------
	// | 1 | 2|  | 6 | 8 |       | 10 | 11 |        | 13 | 15 |   | 18 | 19 |       | 22 | 23 | 25 | 29 |
	// --------  ---------       -----------        ----------    -----------       ---------------------
	tree.Del(30)
	if !checkNodeItems(nil, tree.root, []item{item{12, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{3, 1}, item{9, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{16, 1}, item{21, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[0], []item{item{1, 1}, item{2, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[1], []item{item{6, 1}, item{8, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[0], tree.root.children[0].children[2], []item{item{10, 1}, item{11, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[0], []item{item{13, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[1], []item{item{18, 1}, item{19, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root.children[1], tree.root.children[1].children[2], []item{item{22, 1}, item{23, 1}, item{25, 1}, item{29, 1}}) {
		t.Error("Del")
	}

	//   	           ---------------------------------------------------------------------
	//   	           |     3     |                  11               |    16   |   21    |
	//	               ----------------------------------------------------------------------
	//	              /             \                                  /         |           \
	//         --------              ------------------       -----------   -----------       ---------------------
	//         | 1 | 2|              | 6 | 8 | 9 | 10 |       | 13 | 15 |   | 18 | 19 |       | 22 | 23 | 25 | 29 |
	//         --------              ------------------       -----------   -----------       ---------------------
	tree.Del(12)
	if !checkNodeItems(nil, tree.root, []item{item{3, 1}, item{11, 1}, item{16, 1}, item{21, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[0], []item{item{1, 1}, item{2, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[1], []item{item{6, 1}, item{8, 1}, item{9, 1}, item{10, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[2], []item{item{13, 1}, item{15, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[3], []item{item{18, 1}, item{19, 1}}) {
		t.Error("Del")
	}
	if !checkNodeItems(tree.root, tree.root.children[4], []item{item{22, 1}, item{23, 1}, item{25, 1}, item{29, 1}}) {
		t.Error("Del")
	}

	tree.Del(100)
}

func TestBTreeFind(t *testing.T) {
	size := 1000
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	keys := make([]int, size)
	for i := 0; i < size; i++ {
		keys[i] = r.Int() % (size * 2)
	}

	tree := New(6, intCompare)
	createTree(tree, keys)

	for i := 0; i < size; i++ {
		if _, found := tree.Find(keys[i]); !found {
			t.Error("Find")
		}
	}

	for i := 0; i < size*2; i++ {
		key := keyNotIn(r, keys)
		if _, found := tree.Find(key); found {
			t.Error("Find")
		}
	}
}

type checkNodeInfo struct {
	a      *node
	parent *node
}

func TestBTree(t *testing.T) {
	size := 100000
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	keys := make([]int, size)
	for i := 0; i < size; i++ {
		keys[i] = r.Int() % (size * 2)
	}

	tree := New(7, intCompare)
	createTree(tree, keys)
	l := list.New()
	l.PushBack(checkNodeInfo{tree.root, nil})

	checkRoot := true
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		info := e.Value.(checkNodeInfo)
		if !checkNode(info.parent, info.a, tree.m, checkRoot) {
			t.Error("BTree")
		}

		if checkRoot {
			checkRoot = false
		}

		for i := 0; i < info.a.childCount; i++ {
			l.PushBack(checkNodeInfo{info.a.children[i], info.a})
		}
	}
}

func checkNodeItems(parent *node, a *node, items []item) bool {

	if a.itemLen() != len(items) {
		return false
	}

	if a.children != nil && a.childCount != a.itemLen()+1 {
		return false
	}

	for i := 0; i < len(items); i++ {
		if a.items[i] != items[i] {
			return false
		}
	}

	if a.parent != parent {
		return false
	}

	return true
}

func checkNode(parent *node, a *node, m int, isRoot bool) bool {
	if isRoot {
		if a == nil {
			return true
		}

		if a.itemLen() == 0 {
			return false
		}
	} else {
		if a.itemLen() < m/2 || a.itemLen() == m {
			return false
		}
	}

	if a.children != nil && a.childCount != a.itemLen()+1 {
		return false
	}

	if a.parent != parent {
		return false
	}

	return true
}

func createTree(tree *BTree, keys []int) {
	for i := 0; i < len(keys); i++ {
		tree.Insert(keys[i], 1)
	}
}

func keyNotIn(r *rand.Rand, keys []int) int {
	for {
		key := r.Int()
		exist := false
		for j := 0; j < len(keys); j++ {
			if key == keys[j] {
				exist = true
				break
			}
		}

		if !exist {
			return key
		}
	}
}
