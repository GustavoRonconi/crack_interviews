package btree

// bloomberg interview https://www.youtube.com/watch?v=jma9hFQSCDk&t=78s&ab_channel=KeepOnCoding

type Node struct {
	key    int
	left   *Node
	right  *Node
	parent *Node
}

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
// O(n) complexity and O(1) memory

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
