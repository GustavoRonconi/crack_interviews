package btree

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}

// k = 9
// 20 < 9 = false
// 20 > 9 = true
// n.left == nil = true
// n.left = Node{key: 9}

func (n *Node) insert(k int) {
	if n.key < k {
		if n.right == nil {
			n.right = &Node{key: k, parent: n}
		} else {
			n.right.insert(k)
		}
	} else if n.key > k {
		if n.left == nil {
			n.left = &Node{key: k, parent: n}
		} else {
			n.left.insert(k)
		}
	}
}

// btree
// // // 20 //
// // 9 // 25 //
// 5 // 12 //
// // 11 // 14 //

// first recursive
// root node = 20
// val = 13
// diference = -1
// sucessor = -1

// 20 > 13 = true
// 9 > 13 = false
// diference = 7
// sucessor = 20

// second recursive
// root node = 12
// val = 13
// diference = 7
// sucesseor = 20

// 12 > 13 = false
// 14 > 13 = true
// 14 - 13 = 1
// 1 < 7 = true
// diference = 1
// sucessor = 14

// last recursive
// root node = 14
// val = 13
// diference = 1
// sucessor = 14

//

// func (n *Node) findOrderSucessor(val int, diference int, sucessor int) int {
// 	// greater than value - go to the left
// 	if n.key > val {
// 		if n.left == nil {
// 			return sucessor
// 		}
// 		if n.left.key > val {
// 			if ((n.left.key - val) < diference) || sucessor == -1 {
// 				diference = n.left.key - val
// 				sucessor = n.left.key
// 			}
// 			sucessor = n.left.findOrderSucessor(val, diference, sucessor)
// 		} else {
// 			diference = n.key - val
// 			sucessor = n.key
// 			sucessor = n.left.right.findOrderSucessor(val, diference, sucessor)
// 		}
// 	}

// 	// less than value - go to the right
// 	if n.key < val {
// 		if n.right == nil {
// 			return sucessor
// 		}
// 		if n.right.key > val {
// 			if ((n.right.key - val) < diference) || sucessor == -1 {
// 				diference = n.right.key - val
// 				sucessor = n.right.key
// 			}
// 			sucessor = n.right.findOrderSucessor(val, diference, sucessor)
// 		} else {
// 			diference = n.key - val
// 			sucessor = n.key
// 			sucessor = n.left.right.findOrderSucessor(val, diference, sucessor)
// 		}
// 	}

// 	return sucessor

// }

func (n *Node) findOrderSucessor(val int, diference int, sucessor int) int {
	if n.key > val {
		if ((n.key - val) < diference) || (diference == -1) {
			diference = n.key - val
			sucessor = n.key
		}
	}
	if n.right != nil {
		sucessor = n.right.findOrderSucessor(val, diference, sucessor)
	}

	if n.left != nil {
		sucessor = n.left.findOrderSucessor(val, diference, sucessor)
	}

	return sucessor
}
