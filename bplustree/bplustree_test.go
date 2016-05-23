package bplustree

import "testing"

func TestBplusTreeInsert(t *testing.T) {
	tree := New(3, 3, intCompare)
	var leafs []commonNode

	//  --------
	//  | 5| 8 |
	//  --------
	tree.Insert(5, 1)
	tree.Insert(8, 1)
	if !checkCommonNode(nil, tree.root, []int{5, 8}) {
		t.Error("insert")
	}

	//        -----
	//        | 8 |
	// 	      -----
	//	     /     \
	//  -----      ----------
	//  | 5 |      | 8 | 11 |
	//  -----      ----------
	tree.Insert(11, 1)
	if !checkCommonNode(nil, tree.root, []int{8}) {
		t.Error("insert")
	}
	a := tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{8, 11}) {
		t.Error("insert")
	}

	leafs = append(leafs, a.children[0], a.children[1])
	if !checkLeafList(leafs) {
		t.Error("insert")
	}

	//        ----------
	//        | 8 | 11 |
	// 	      ----------
	//	     /    |    \
	//  -----   -----   -----------
	//  | 5 |   | 8 |   | 11 | 15 |
	//  -----   -----   -----------
	tree.Insert(15, 1)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{8, 11}) {
		t.Error("insert")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{8}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[2], []int{11, 15}) {
		t.Error("insert")
	}

	leafs = append(leafs, a.children[0], a.children[1], a.children[2])
	if !checkLeafList(leafs) {
		t.Error("insert")
	}

	//                --------
	//                |  11  |
	//                --------
	//			     /        \
	//        -------          --------
	//        |  8  |          |  15  |
	// 	      -------          --------
	//	     /      \         /        \
	//   -----      -----   ------    -----------
	//   | 5 |      | 8 |   | 11 |    | 15 | 20 |
	//   -----      -----   ------    -----------
	tree.Insert(20, 1)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("insert")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{8}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{15}) {
		t.Error("insert")
	}

	child := a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{5}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{8}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{11}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{15, 20}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("insert")
	}

	//                       --------
	//                       |  11  |
	//                       --------
	//			            /        \
	//     -----------------          --------
	//     |   5   |   8   |          |  15  |
	// 	   -----------------          --------
	//	  /        |        \         /        \
	// -----   ---------    -----   ------    -----------
	// | 4 |   | 5 | 6 |    | 8 |   | 11 |    | 15 | 20 |
	// -----   ---------    -----   ------    -----------
	tree.Insert(4, 1)
	tree.Insert(6, 1)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("insert")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5, 8}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{15}) {
		t.Error("insert")
	}

	child = a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{5, 6}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[2], []int{8}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1], child.children[2])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{11}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{15, 20}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("insert")
	}

	//         ----------------------------------------------
	//         |           8          |           11        |
	//         ----------------------------------------------
	//        /                       |                     \
	//     ---------              ---------                  --------
	//     |   5   |              |   9   |                  |  15  |
	// 	   ---------              ---------                  --------
	//	  /         \             /        \                /        \
	// -----     ---------     -----    ----------        ------    -----------
	// | 4 |     | 5 | 6 |     | 8 |    | 9 | 10 |        | 11 |    | 15 | 20 |
	// -----     ---------     -----    ----------        ------    -----------
	tree.Insert(9, 1)
	tree.Insert(10, 1)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{8, 11}) {
		t.Error("insert")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{9}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[2], []int{15}) {
		t.Error("insert")
	}

	child = a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{5, 6}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{8}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{9, 10}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[2].(*node)
	if !checkCommonNode(child, child.children[0], []int{11}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{15, 20}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("insert")
	}

	//                                 ---------------------
	//                                 |       11          |
	//                                 ---------------------
	//                                /                     \
	//         ------------------------                     ----------------------
	//         |           8          |                     |         15         |
	//         ------------------------                     ----------------------
	//        /                       \                    /                      \
	//     ---------               ---------            ----------                 --------
	//     |   5   |               |   9   |            |   12   |                 |  20  |
	// 	   ---------               ---------            ----------                 --------
	//	  /         \              /        \           /        \                /        \
	// -----     ---------     -----    ----------     ------    -----------   ------    -----------
	// | 4 |     | 5 | 6 |     | 8 |    | 9 | 10 |     | 11 |    | 12 | 13 |   | 15 |    | 20 | 21 |
	// -----     ---------     -----    ----------     ------    -----------   ------    -----------
	tree.Insert(12, 1)
	tree.Insert(13, 1)
	tree.Insert(21, 1)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("insert")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{8}) {
		t.Error("insert")
	}
	if !checkCommonNode(a, a.children[1], []int{15}) {
		t.Error("insert")
	}

	child = a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{5}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{9}) {
		t.Error("insert")
	}

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{12}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{20}) {
		t.Error("insert")
	}

	child = a.children[0].(*node).children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{5, 6}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[0].(*node).children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{8}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{9, 10}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[1].(*node).children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{11}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{12, 13}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[1].(*node).children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{15}) {
		t.Error("insert")
	}
	if !checkCommonNode(child, child.children[1], []int{20, 21}) {
		t.Error("insert")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("insert")
	}
}

func TestBplusTreeDel(t *testing.T) {
	tree := New(3, 3, intCompare)
	var leafs []commonNode

	//                       --------
	//                       |  11  |
	//                       --------
	//			            /        \
	//     -----------------          --------
	//     |   5   |   8   |          |  15  |
	// 	   -----------------          --------
	//	  /        |        \         /        \
	// -----   ---------    -----   ------    -----------
	// | 4 |   | 5 | 6 |    | 8 |   | 11 |    | 15 | 20 |
	// -----   ---------    -----   ------    -----------
	tree.Insert(5, 1)
	tree.Insert(8, 1)
	tree.Insert(11, 1)
	tree.Insert(15, 1)
	tree.Insert(20, 1)
	tree.Insert(4, 1)
	tree.Insert(6, 1)

	//                       --------
	//                       |  11  |
	//                       --------
	//			            /        \
	//     -----------------          --------
	//     |   5   |   8   |          |  20  |
	// 	   -----------------          --------
	//	  /        |        \         /        \
	// -----   ---------    -----   ------    ------
	// | 4 |   | 5 | 6 |    | 8 |   | 15 |    | 20 |
	// -----   ---------    -----   ------    ------
	tree.Del(11)
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("del")
	}

	a := tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5, 8}) {
		t.Error("del")
	}
	if !checkCommonNode(a, a.children[1], []int{20}) {
		t.Error("del")
	}

	child := a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{5, 6}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[2], []int{8}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1], child.children[2])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{15}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{20}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("del")
	}

	//                       --------
	//                       |  11  |
	//                       --------
	//			            /        \
	//     -----------------          --------
	//     |   5   |   6   |          |  20  |
	// 	   -----------------          --------
	//	  /        |        \         /        \
	// -----     -----    -----   ------    ------
	// | 4 |     | 5 |    | 6 |   | 15 |    | 20 |
	// -----     -----    -----   ------    ------
	tree.Del(8)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("del")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5, 6}) {
		t.Error("del")
	}
	if !checkCommonNode(a, a.children[1], []int{20}) {
		t.Error("del")
	}

	child = a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{5}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[2], []int{6}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1], child.children[2])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{15}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{20}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("del")
	}

	//                       --------
	//                       |  11  |
	//                       --------
	//			            /        \
	//             ---------          --------
	//             |   5   |          |  20  |
	// 	           ---------          --------
	//	          /         \         /       \
	//       -----          -----   ------    ------
	//       | 4 |          | 5 |   | 15 |    | 20 |
	//       -----          -----   ------    ------
	tree.Del(6)
	leafs = leafs[0:0]
	if !checkCommonNode(nil, tree.root, []int{11}) {
		t.Error("del")
	}

	a = tree.root.(*node)
	if !checkCommonNode(a, a.children[0], []int{5}) {
		t.Error("del")
	}
	if !checkCommonNode(a, a.children[1], []int{20}) {
		t.Error("del")
	}

	child = a.children[0].(*node)
	if !checkCommonNode(child, child.children[0], []int{4}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{5}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	child = a.children[1].(*node)
	if !checkCommonNode(child, child.children[0], []int{15}) {
		t.Error("del")
	}
	if !checkCommonNode(child, child.children[1], []int{20}) {
		t.Error("del")
	}
	leafs = append(leafs, child.children[0], child.children[1])

	if !checkLeafList(leafs) {
		t.Error("del")
	}
}

func checkCommonNode(parent *node, a commonNode, keys []int) bool {
	if a.size() != len(keys) {
		return false
	}

	for i := 0; i < len(keys); i++ {
		if a.getKey(i) != keys[i] {
			return false
		}
	}

	if a.getParent() != parent {
		return false
	}

	return true
}

func checkLeafList(leafs []commonNode) bool {
	if leafs[0].(*leaf).pre != nil {
		return false
	}

	if leafs[len(leafs)-1].(*leaf).next != nil {
		return false
	}

	for i := 0; i < len(leafs)-1; i++ {
		if leafs[i].(*leaf).next != leafs[i+1] {
			return false
		}

		if leafs[i+1].(*leaf).pre != leafs[i] {
			return false
		}
	}

	return true
}
