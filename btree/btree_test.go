package btree

// test btree

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBinaryTree(t *testing.T) {
	binaryTree := &Node{key: 20, left: nil, right: nil}

	binaryTree.insert(20)
	binaryTree.insert(9)
	binaryTree.insert(5)
	binaryTree.insert(12)
	binaryTree.insert(11)
	binaryTree.insert(14)
	binaryTree.insert(25)

	assert.Equal(t, binaryTree.findOrderSucessor(50, -1, -1), -1)
	assert.Equal(t, binaryTree.findOrderSucessor(9, -1, -1), 11)
	assert.Equal(t, binaryTree.findOrderSucessor(14, -1, -1), 20)

}
